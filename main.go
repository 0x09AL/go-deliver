package main
import (
	"go-deliver/terminal"
	"go-deliver/database"
	"go-deliver/model"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	_ "database/sql"
	_"github.com/mattn/go-sqlite3"
	"gopkg.in/gcfg.v1"
	"fmt"
)


func main() {

	Configuration := model.CFG{}
	// Setup Routes
	listener := mux.NewRouter()
	listener.HandleFunc("/",database.ShowShit).Methods("GET")
	listener.HandleFunc("/{guid}/",database.GetPayload).Methods("GET")


	err := gcfg.ReadFileInto(&Configuration,"config.conf")

	if err != nil {
		log.Fatalf("Failed to parse gcfg data: %s", err)
	}
	log.Println(fmt.Sprintf("Starting http server on port %d",Configuration.Http.Port))

	// Starts the server
	go http.ListenAndServe(fmt.Sprintf(":%d",Configuration.Http.Port),listener)

	// Start the terminal
	terminal.StartTerminal()


}