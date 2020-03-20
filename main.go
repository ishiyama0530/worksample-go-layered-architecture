package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"worksample-go-layered-architecture/app/handler/v1"
	apiconfig "worksample-go-layered-architecture/config"
	"github.com/julienschmidt/httprouter"
)

const env = "dev"

func main() {
	apiconfig.Start(env)
	config := apiconfig.Get()

	var router *httprouter.Router = httprouter.New()
	handler.RegisterHandleFunc(router)

	server := http.Server{
		Addr:    config.Server.Addr,
		Handler: router,
	}

	fmt.Printf("Listening at: %s\n", server.Addr)
	server.ListenAndServe()
}
