package main

import (
	"database/sql"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"complete backend/services/user"
)


type server struct{
	addr string
	db *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *server{
	return &server{
		addr: addr,
		db: db,
	}
}


func (se *server) RUN() error{
	router :=  mux.NewRouter()
	subrouter := router.PathPerifx("api/v1").Subrouter()

	userhandler := user.newhandler()
	userhandler.RegisterRoutes(subrouter)

	fmt.Println("server listen on",se.addr)
	return http.ListenAndServe(se.addr,router)
}