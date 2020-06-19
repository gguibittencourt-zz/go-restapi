package main

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/gguibittencourt/go-restapi/config"
	. "github.com/gguibittencourt/go-restapi/config/dao"
	userRouter "github.com/gguibittencourt/go-restapi/router"
	"github.com/gorilla/mux"
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

	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func handleUserRouter(router *mux.Router) {
	router.HandleFunc(userPath, userRouter.List).Methods("GET")
	router.HandleFunc(userPathId, userRouter.GetByID).Methods("GET")
	router.HandleFunc(userPath, userRouter.Create).Methods("POST")
	router.HandleFunc(userPathId, userRouter.Update).Methods("PUT")
	router.HandleFunc(userPathId, userRouter.Delete).Methods("DELETE")
}
