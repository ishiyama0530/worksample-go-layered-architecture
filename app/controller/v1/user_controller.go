package v1

import (
	"fmt"
	"net/http"
	uc "worksample-go-layered-architecture/app/usecase/user"

	"github.com/julienschmidt/httprouter"
)

// UserController is
type UserController struct {
	interacor uc.GetUsecase
}

// NewUserController is
func NewUserController(interacor uc.GetUsecase, router *httprouter.Router) *UserController {
	c := new(UserController)
	c.interacor = interacor

	router.GET("/account/:id", c.get)

	return c
}

func (c *UserController) get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	inputData := uc.NewGetInputData(id)
	outputData, err := c.interacor.Handle(*inputData)
	if err != nil {
		fmt.Fprintf(w, err.Error())

	} else {
		fmt.Fprintf(w, "hello, %s\n", outputData.Name)

	}
}
