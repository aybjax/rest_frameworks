package main

import (
	"database/sql"
	"go-restful/dbutils"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// connect db
	var err error
	DB, err = sql.Open("sqlite3", "./railapi.db")

	if err != nil {
		log.Fatal("Driver creation failed")
	}

	dbutils.Initialize(DB)

	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	
	t := Train{}
	t.Register(wsContainer)

	log.Println("Start listening on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}

	log.Fatal(server.ListenAndServe())
}