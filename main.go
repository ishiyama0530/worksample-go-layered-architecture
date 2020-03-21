package main

import (
	"fmt"
	"net/http"
	"worksample-go-layered-architecture/inject"

	config "worksample-go-layered-architecture/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

const env = "mock"

func main() {
	config.Start(env)
	c := config.Get()

	var router *httprouter.Router = httprouter.New()

	inject.Setup(router)

	server := http.Server{
		Addr:    c.Server.Addr,
		Handler: router,
	}

	fmt.Printf("Listening at: %s\n", server.Addr)
	server.ListenAndServe()
}
