package main

import (
	"fmt"
	. "github.com/gguibittencourt/go-restapi/config"
	. "github.com/gguibittencourt/go-restapi/config/dao"
	. "github.com/gguibittencourt/go-restapi/router"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var dao = UsersDAO{}
var config = Config{}

const userPath = "/api/users"
const userPathId = userPath + "/{id}"
const port = ":3000"

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	router := mux.NewRouter()
	handleUserRouter(router)
	handleFileUploadRouter(router)

	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS()(router)))
}

func handleUserRouter(router *mux.Router) {
	router.HandleFunc(userPath, List).Methods("GET")
	router.HandleFunc(userPathId, GetByID).Methods("GET")
	router.HandleFunc(userPath, Create).Methods("POST")
	router.HandleFunc(userPathId, Update).Methods("PUT")
	router.HandleFunc(userPathId, Delete).Methods("DELETE")
}

func handleFileUploadRouter(router *mux.Router) {
	router.HandleFunc("/api/file-upload", FileUpload).Methods("POST")
}
