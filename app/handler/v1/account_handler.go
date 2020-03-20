package handler

import (
	"fmt"
	"net/http"
	ap "worksample-go-layered-architecture/app/application/user"
	"worksample-go-layered-architecture/app/domain/user"
	"worksample-go-layered-architecture/app/infrastructure/inmemory"
	uc "worksample-go-layered-architecture/app/usecase/user"

	"github.com/julienschmidt/httprouter"
)

// RegisterHandleFunc is
func RegisterHandleFunc(router *httprouter.Router) {
	router.GET("/account/:id", get)
}

func get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	var repo user.Repository = &inmemory.UserRepository{}
	var interacor uc.GetUsecase = ap.NewGetInteractor(repo)
	inputData := uc.NewGetInputData(id)
	outputData, err := interacor.Handle(*inputData)
	if err != nil {
		fmt.Fprintf(w, err.Error())

	} else {
		fmt.Fprintf(w, "hello, %s\n", outputData.Name)

	}
}
