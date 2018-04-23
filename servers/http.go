package servers

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"go-deliver/database"
	"go-deliver/model"
	"io/ioutil"
)





func StartHTTPListener(Configuration model.CFG)  {

	Config = Configuration
	listener := mux.NewRouter()
	listener.HandleFunc("/",database.ShowIndex).Methods("GET")
	listener.HandleFunc("/{guid}/",database.GetPayload).Methods("GET")
	listener.NotFoundHandler = http.HandlerFunc(handle404Error)

	server := &http.Server{
		Addr:fmt.Sprintf(":%d",Config.Http.Port),
		Handler:listener,
	}


	server.ListenAndServe()
}

func handle404Error(w http.ResponseWriter, r *http.Request){

	bData, err := ioutil.ReadFile(Config.Http.Template404)
	if err != nil {
		w.Write([]byte("What are you lookin for ?"))

	}
	w.Write(bData)


}
