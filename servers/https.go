package servers

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"go-deliver/database"
	"go-deliver/model"
	"io/ioutil"
	"log"
)





func StartHTTPSListener(Configuration model.CFG)  {

	Config = Configuration
	listener := mux.NewRouter()
	listener.HandleFunc("/",database.ShowIndex).Methods("GET")
	listener.HandleFunc("/{guid}/",database.GetPayload).Methods("GET")
	listener.NotFoundHandler = http.HandlerFunc(handle404ErrorSSL)

	err := http.ListenAndServeTLS(fmt.Sprintf(":%d",Config.Https.Port), Config.Https.Publickey, Config.Https.Privatekey, listener)
	if err != nil {
		log.Fatal("Error starting HTTPS Server : ", err)
	}

}

func handle404ErrorSSL(w http.ResponseWriter, r *http.Request){

	bData, err := ioutil.ReadFile(Config.Https.Template404)
	if err != nil {
		w.Write([]byte("What are you lookin for ?"))

	}
	w.Write(bData)

}
