package main

import (
	"fmt"
	"net/http"

	v1 "worksample-go-layered-architecture/app/controller/v1"
	config "worksample-go-layered-architecture/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

const env = "dev"

func main() {
	config.Start(env)
	c := config.Get()

	var router *httprouter.Router = httprouter.New()
	v1.NewAccountController(router)

	server := http.Server{
		Addr:    c.Server.Addr,
		Handler: router,
	}

	fmt.Printf("Listening at: %s\n", server.Addr)
	server.ListenAndServe()
}
