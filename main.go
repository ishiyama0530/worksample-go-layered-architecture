package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"worksample-go-layered-architecture/inject"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

const env = "mock"

func main() {
	log.Println(fmt.Sprintf("./config/%s.env", os.Getenv("GO_ENV")))
	err := godotenv.Load(fmt.Sprintf("./config/%s.env", os.Getenv("GO_ENV")))

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var router *httprouter.Router = httprouter.New()
	inject.Setup(router)

	server := http.Server{
		Addr:    os.Getenv("IP"),
		Handler: router,
	}

	log.Printf("Listening at: %s\n", server.Addr)
	server.ListenAndServe()
}
