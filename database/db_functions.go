package database


import (
	"net/http"
	"fmt"
	_ "database/sql"
	_"github.com/mattn/go-sqlite3"
	"database/sql"
	"log"
	"html/template"
	"github.com/gorilla/mux"
	"net"
	"encoding/base64"
	"io/ioutil"
	"go-deliver/model"

	"math/rand"
	"time"
)



var db, _ = sql.Open("sqlite3", "test_db.db")
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



	// This function will create the requrired shits
	createTableSql := `CREATE TABLE payloads (
							id	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
							name	TEXT NOT NULL UNIQUE,
							content_type	TEXT,
							host_blacklist	TEXT,
							host_whitelist	TEXT,
							data_file	TEXT,
							data_b64	TEXT,
							type_id	INTEGER NOT NULL,
							guid	TEXT NOT NULL UNIQUE
						);
						`

	createHostSql := `CREATE TABLE hosts (
							id	INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
							type	TEXT NOT NULL,
							data	TEXT NOT NULL
						);`

	createTypesSql := `CREATE TABLE types (
							id	INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
							name	TEXT NOT NULL UNIQUE,
							type_template	TEXT
						);`



	fmt.Println(createTableSql)
	fmt.Println(createHostSql)
	fmt.Println(createTypesSql)
}


func EditPayload(w http.ResponseWriter,r *http.Request)  {
	vars := mux.Vars(r)
	pid := vars["pid"]
	ReadPayload := "SELECT * FROM payloads WHERE p_id=?"
	rows, err := db.Query(ReadPayload, pid)
	if err != nil {
		panic(err)
	}
	payload := model.Payload{}
	rows.Next()

	err_sql := rows.Scan(&payload.Id,&payload.Name,&payload.Content_type,&payload.Data_b64)

	if err_sql != nil{
		panic(err_sql)
	}
	t,err_tmpl := template.ParseFiles("templates/edit.html")
	if err_tmpl != nil{
		panic(err_tmpl)
	}
	t.Execute(w,payload)
}
func DeletePayload(w http.ResponseWriter,r *http.Request)  {

}


func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func GetTypeid(ptype string) int{
	// To be Implemented
	return 1;
}


func InsertPayload(payload model.Payload){
	randomInit()
	query := "INSERT INTO payloads VALUES (NULL,?,?,?,?,?,?,?,?);"
	tx, _ := db.Begin()
	stmt, err_stmt := tx.Prepare(query)
	payload.Guid = RandStringRunes(32)

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
func GetPayloads(w http.ResponseWriter,r *http.Request)  {


}
func GetPayload(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	guid := vars["guid"]
	GetPayloadQuery := `SELECT id,
						name,
						content_type,
						COALESCE(host_blacklist, '') as host_blacklist, 
						COALESCE(host_whitelist, '') as host_whitelist,
						COALESCE(data_file, '') as data_file, 
						COALESCE(data_b64, '') as data_b64 ,
						type_id 
						from payloads 
						WHERE guid=?`


	rows, err := db.Query(GetPayloadQuery, guid)
	if err != nil {
		panic(err)
	}
	payload := model.Payload{}
	rows.Next()
	err_sql := rows.Scan(&payload.Id,&payload.Name,&payload.Content_type,&payload.Host_blacklist,&payload.Host_whitelist,&payload.Data_file,&payload.Data_b64,&payload.Type_id)

	if err_sql != nil{
		panic(err_sql)
	}

	ip , _ , _ := net.SplitHostPort(r.RemoteAddr)

	log.Println(fmt.Sprintf("Delivering payload %s to IP : %s",payload.Name,ip))

	w.Header().Set("Content-Type",payload.Content_type)
	w.WriteHeader(http.StatusOK)

	if payload.Data_file == ""{
		if payload.Data_b64 != ""{
			data, err := base64.StdEncoding.DecodeString(payload.Data_b64)
			if err != nil{
				log.Println("ERROR : Decoding b64 payload failed.")
			}
			w.Write([]byte(data))
		}else{
			log.Println("ERROR : Payload delivery failed. No content or file specified.")
		}
	}else{
		// Write data from file
		data, err := ioutil.ReadFile(payload.Data_file)
		if err != nil{
			log.Println(fmt.Sprintf("ERROR: Payload file %s not found.", payload.Data_file))
			return
		}
		w.Write(data)
	}


}


func CreatePayloadGet(w http.ResponseWriter,r *http.Request)  {
	w.WriteHeader(http.StatusOK)
	t, err_tmpl := template.ParseFiles("templates/create.html")
	if err_tmpl != nil{
		panic(err_tmpl)
	}
	t.Execute(w,nil)
}



func CreatePayload(w http.ResponseWriter,r *http.Request)  {

	tx, _ := db.Begin()
	stmt, err_stmt := tx.Prepare("INSERT INTO payloads VALUES (NULL,?,?,?);")
	if err_stmt != nil {
		log.Fatal(err_stmt)
	}
	_, err := stmt.Exec(r.FormValue("p_name"),r.FormValue("p_ct"),r.FormValue("p_content"))
	tx.Commit()
	if err != nil {
		log.Println("Payload Insertion Failed")
		w.WriteHeader(http.StatusBadRequest)
	}else{
		http.Redirect(w,r,"/create.html",http.StatusSeeOther)
		log.Println("Payload created.")


	}

}