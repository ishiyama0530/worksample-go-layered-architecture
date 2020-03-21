package v1

import (
	"fmt"
	"net/http"
	ap "worksample-go-layered-architecture/app/application/user"
	"worksample-go-layered-architecture/app/infrastructure/inmemory"
	uc "worksample-go-layered-architecture/app/usecase/user"

	"github.com/julienschmidt/httprouter"
)

// AccountController is
type AccountController struct {
	interacor uc.GetUsecase
}

// NewAccountController is
func NewAccountController(router *httprouter.Router) *AccountController {
	c := new(AccountController)
	repo := &inmemory.UserRepository{}
	c.interacor = ap.NewGetInteractor(repo)

	router.GET("/account/:id", c.get)

	return c
}

func (c *AccountController) get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	inputData := uc.NewGetInputData(id)
	outputData, err := c.interacor.Handle(*inputData)
	if err != nil {
		fmt.Fprintf(w, err.Error())

	} else {
		fmt.Fprintf(w, "hello, %s\n", outputData.Name)

	}
}
