package v1

import (
	"fmt"
	"net/http"
	uc "worksample-go-layered-architecture/app/usecase/user"

	"github.com/julienschmidt/httprouter"
)

// UserController is
type UserController struct {
	getInteractor    uc.GetUsecase
	createInteractor uc.CreateUsecase
}

// NewUserController is
func NewUserController(
	getInteractor uc.GetUsecase,
	createInteractor uc.CreateUsecase,
	router *httprouter.Router) *UserController {
	c := new(UserController)

	c.getInteractor = getInteractor
	c.createInteractor = createInteractor

	router.GET("/user/:id", c.get)
	// TODO router.POST("/user/:name", c.get)

	return c
}

func (c *UserController) get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	inputData := uc.NewGetInputData(id)
	outputData, err := c.getInteractor.Handle(*inputData)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		fmt.Fprintf(w, "get: %s\n", outputData.Name)
	}
}

func (c *UserController) post(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("name")
	inputData := uc.NewCreateInputData(name)
	outputData, err := c.createInteractor.Handle(*inputData)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		fmt.Fprintf(w, "post: %s\n", outputData.Name)
	}
}
