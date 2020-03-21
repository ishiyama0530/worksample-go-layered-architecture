package v1

import (
	"fmt"
	"net/http"
	uc "worksample-go-layered-architecture/app/usecase/user"

	"github.com/julienschmidt/httprouter"
)

// AccountController is
type AccountController struct {
	interacor uc.GetUsecase
}

// Setup is
func Setup(interacor uc.GetUsecase, router *httprouter.Router) {
	c := new(AccountController)
	c.interacor = interacor

	router.GET("/account/:id", c.get)
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
