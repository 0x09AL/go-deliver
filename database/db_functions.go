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
)



var db, _ = sql.Open("sqlite3", "db_file.db")
type Payload struct {
	id int
	name string
	content_type string
	host_blacklist	string
	host_whitelist	string
	data_file	string
	data_b64	string
	ptype 		string
	one_liner	string

}

type Host struct {
	name string
	htype string
	data string
}

func CreateTable()  {


	// This function will create the requrired shits
	createTableSql := `CREATE TABLE payloads (
							id	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
							name	TEXT NOT NULL UNIQUE,
							content_type	TEXT,
							host_blacklist	TEXT,
							host_whitelist	TEXT,
							data_file	TEXT,
							data_b64	TEXT,
							type	INTEGER NOT NULL
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
	payload := Payload{}
	rows.Next()
	err_sql := rows.Scan(&payload.id,&payload.name,&payload.content_type,&payload.data_b64)

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
func ShowShit(w http.ResponseWriter,r *http.Request)  {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w,"What you received is shit")
}
func GetPayloads(w http.ResponseWriter,r *http.Request)  {

}
func GetPayload(w http.ResponseWriter,r *http.Request){

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