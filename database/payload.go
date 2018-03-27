package database


import (
	"net/http"
	"fmt"
	_ "database/sql"
	_"github.com/mattn/go-sqlite3"
	"log"
	"github.com/gorilla/mux"
	"net"
	"encoding/base64"
	"io/ioutil"
	"go-deliver/model"
	"math/rand"
	"time"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)




var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")


func randomInit() {
	// There is no reason to use a strong number generator since the payloads already have whitelist/blacklists.
	rand.Seed(time.Now().UnixNano())
}



func CreateTable()  {

	// Insert payload query
	//insert into payloads values (NULL,'test123','wtf',NULL,NULL,'/tmp/shit',NULL,1);
	// Payload types query
	//insert into payload_types values (NULL,'javascript',NULL)



	// This function will create the requrired databases


}


func DeletePayload(name string)  {

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare(model.DeletePayloadQuery)
	_, err := stmt.Exec(name)
	if err != nil{
		log.Panic(err)
	}
	tx.Commit()
	log.Println(fmt.Sprintf("Success : Payload %s deleted .",name))
	}


func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func GetTypeid(ptype string) (int , string) {

	payloadType := model.PayloadType{}
	err := db.QueryRow(model.GetPayloadTypeId, ptype).Scan(&payloadType.Type_id,&payloadType.Content_type)
	if err != nil {

		fmt.Println(fmt.Sprintf("ERROR: Payload type %s doesn't exist.",ptype))
		return 0,""
	}
	return payloadType.Type_id, payloadType.Content_type;
}


func InsertPayload(payload model.Payload){
	randomInit()

	tx, _ := db.Begin()
	stmt, err_stmt := tx.Prepare(model.InsertPayloadQuery)
	payload.Guid = RandStringRunes(32)

	fmt.Println("")
	if err_stmt != nil {
		log.Fatal(err_stmt)
	}
	_, err := stmt.Exec(payload.Name,payload.Content_type,payload.Host_blacklist,payload.Host_whitelist,payload.Data_file,payload.Data_b64,payload.Type_id,payload.Guid)
	tx.Commit()
	if err != nil{
		log.Println("ERROR: Error inserting payload.")
	}else{

		log.Println(fmt.Sprintf("Payload with name %s created successfully.",payload.Name)) // Fix the URL output
	}

}


func ShowShit(w http.ResponseWriter,r *http.Request)  {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w,"What you received is shit")
}

func GetPayloads()  {

	rows, err := db.Query(model.GetPayloadsQuery)

	payload := model.Payload{}
	if err != nil {
		panic(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Name", "Path", "Content Type", "Hosts Blacklist" , "Hosts Whitelist"})

	for rows.Next() {
		err := rows.Scan(&payload.Id, &payload.Name, &payload.Guid, &payload.Content_type,&payload.Host_whitelist, &payload.Host_blacklist)
		table.Append([]string{ strconv.Itoa(payload.Id),payload.Name,fmt.Sprintf("/%s/",payload.Guid),payload.Content_type,payload.Host_whitelist,payload.Host_blacklist})
		if err != nil {
			log.Fatal(err)
		}
	}
	table.Render()
}


func GetPayload(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	guid := vars["guid"]
	var response []byte
	allow := false

	payload := model.Payload{}
	err := db.QueryRow(model.GetPayloadQuery, guid).Scan(&payload.Id,&payload.Name,&payload.Content_type,&payload.Host_blacklist,&payload.Host_whitelist,&payload.Data_file,&payload.Data_b64,&payload.Type_id)
	if err != nil {
		panic(err)
	}

	ip , _ , _ := net.SplitHostPort(r.RemoteAddr)



	if payload.Data_file == ""{
		if payload.Data_b64 != ""{
			data, err := base64.StdEncoding.DecodeString(payload.Data_b64)
			if err != nil{
				log.Println("ERROR : Decoding b64 payload failed.")
				return
			}else{
				response = data
			}

		}else{
			log.Println("ERROR : Payload delivery failed. No content or file specified.")
			response = []byte("")
		}
	}else{
		// Write data from file
		data, err := ioutil.ReadFile(payload.Data_file)
		if err != nil{
			log.Println(fmt.Sprintf("ERROR: Payload file %s not found.", payload.Data_file))
			response = []byte("")
		}
		response = data
	}


	if payload.Host_whitelist != ""{
		htype , data := GetData(payload.Host_whitelist)
		switch htype{
		case "ip":
			//Check if remote ip is the same
			if data == ip{
				allow = true
			}

		case "subnet":
			// Check if ip is in subnet
			_ , ipNet , _ := net.ParseCIDR(data)
			if ipNet.Contains(net.ParseIP(ip)){
				allow = true
			}

		}

	}else if payload.Host_blacklist != ""{
		htype , data := GetData(payload.Host_blacklist)
		switch htype{
			case "ip":
				//Check if remote ip is the same
				if data == ip{
					allow = false
				}else {
					allow = true
				}
			case "subnet":
				// Check if ip is in subnet
				_ , ipNet , _ := net.ParseCIDR(data)
				if ipNet.Contains(net.ParseIP(ip)){
					allow = false
				}else{
					allow = true
					}
		}
	}

	if allow == true{
		fmt.Println("") // Prints a new line
		log.Println(fmt.Sprintf("Delivering payload %s to IP : %s",payload.Name,ip))

		w.Header().Set("Content-Type",payload.Content_type)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}else{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("")) // Will write the 404 template later.
		log.Println(fmt.Sprintf("Denided access to IP: %s for %s payload.",ip,payload.Name))
	}


}

func GetPayloadTypes() []model.PayloadType{

	payloadTypes := []model.PayloadType{}

	rows, err := db.Query(model.GetPayloadTypesQuery)

	payloadType := model.PayloadType{}

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err := rows.Scan(&payloadType.Type_name,&payloadType.Content_type)
		payloadTypes = append(payloadTypes,payloadType)
		if err != nil {
			log.Fatal(err)
		}
	}

	return payloadTypes
}