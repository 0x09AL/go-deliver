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

func FuckOFF()  {
	createTableSql := `CREATE TABLE payloads (
p_id	INTEGER PRIMARY KEY AUTOINCREMENT,
p_name	TEXT UNIQUE,
p_ct	TEXT NOT NULL,
p_content	TEXT NOT NULL
);
`
	fmt.Println(createTableSql)
}

type Payload struct{
	P_id int
	P_name	string
	P_ct	string
	P_content	string
}

var db, _ = sql.Open("sqlite3", "db_file.db")


func CreateTable()  {
	createTableSql := `CREATE TABLE payloads (
						p_id	INTEGER PRIMARY KEY AUTOINCREMENT,
						p_name	TEXT UNIQUE,
						p_ct	TEXT NOT NULL,
						p_content	TEXT NOT NULL
						);
						`
						// This function will create the table if it doesn't exists.
	fmt.Println(createTableSql)
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
	err_sql := rows.Scan(&payload.P_id,&payload.P_name,&payload.P_ct,&payload.P_content)

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