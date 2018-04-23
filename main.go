package main
import (
	"go-deliver/terminal"
	"go-deliver/model"
	"go-deliver/servers"
	"log"
	_ "database/sql"
	_"github.com/mattn/go-sqlite3"
	"gopkg.in/gcfg.v1"
	"fmt"
)


func main() {

	Configuration := model.CFG{}

	err := gcfg.ReadFileInto(&Configuration,"config.conf")

	if err != nil {
		log.Fatalf("Failed to parse gcfg data: %s", err)
	}

	if(Configuration.Http.Enable == "true"){
		log.Println(fmt.Sprintf("Starting http server on port %d",Configuration.Http.Port))
		go servers.StartHTTPListener(Configuration)
	}
	if(Configuration.Https.Enable == "true"){
		log.Println(fmt.Sprintf("Starting https server on port %d",Configuration.Https.Port))
		go servers.StartHTTPSListener(Configuration)
	}


	// Start the terminal
	terminal.StartTerminal()


}