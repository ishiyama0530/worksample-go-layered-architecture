package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ishiyama0530/worksample-go-layered-architecture/app/infra/inmemory"
	usecase "github.com/ishiyama0530/worksample-go-layered-architecture/app/usecase/user"

	"github.com/julienschmidt/httprouter"
)

// RegisterHandleFunc is
func RegisterHandleFunc(router *httprouter.Router) {
	router.GET("/account/:id", get)
}

func get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	store := &inmemory.UserStore{}
	interactor := usecase.NewGetUserInteractor(store)
	user, _ := interactor.GetUser(id)

	fmt.Fprintf(w, "hello, %s\n", user.Name)
}
