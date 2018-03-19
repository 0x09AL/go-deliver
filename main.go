package main
import (
	"go-deliver/terminal"
	"go-deliver/database"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	_ "database/sql"
	_"github.com/mattn/go-sqlite3"
)


func main() {


	// Setup Routes
	listener := mux.NewRouter()
	listener.HandleFunc("/",database.ShowShit).Methods("GET")
	listener.HandleFunc("/{puid}",database.GetPayload).Methods("GET")
	listener.HandleFunc("/payloads/edit/{pid}",database.EditPayload).Methods("GET")
	listener.HandleFunc("/payloads/new",database.CreatePayloadGet).Methods("GET")
	listener.HandleFunc("/payloads/create",database.CreatePayload).Methods("POST")

	// Starts the server
	log.Println("Starting server .")
	go http.ListenAndServe(":8000",listener)

	// Start the terminal
	terminal.StartTerminal()


}