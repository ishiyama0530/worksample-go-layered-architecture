package main

import (
	"fmt"
	"net/http"

	ap "worksample-go-layered-architecture/app/application/user"
	v1 "worksample-go-layered-architecture/app/controller/v1"
	infra "worksample-go-layered-architecture/app/infrastructure/inmemory"
	config "worksample-go-layered-architecture/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

const env = "dev"

func main() {
	config.Start(env)
	c := config.Get()

	var router *httprouter.Router = httprouter.New()

	userRepo := infra.NewUserRepository()
	getUserInteractor := ap.NewGetInteractor(userRepo)

	v1.Setup(getUserInteractor, router)

	server := http.Server{
		Addr:    c.Server.Addr,
		Handler: router,
	}

	fmt.Printf("Listening at: %s\n", server.Addr)
	server.ListenAndServe()
}
