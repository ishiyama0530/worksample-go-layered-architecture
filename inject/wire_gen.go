// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package inject

import (
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	user2 "worksample-go-layered-architecture/app/application/user"
	"worksample-go-layered-architecture/app/controller/v1"
	user3 "worksample-go-layered-architecture/app/domain/user"
	"worksample-go-layered-architecture/app/persistence/inmemory"
	"worksample-go-layered-architecture/app/queryservice/inmemory/user"
	user4 "worksample-go-layered-architecture/app/usecase/user"
)

// Injectors from Injector.mock.go:

func Setup(router *httprouter.Router) *v1.UserController {
	getInteractor := user.NewGetInteractor()
	userRepository := inmemory.NewUserRepository()
	createInteractor := user2.NewCreateInteractor(userRepository)
	userController := v1.NewUserController(getInteractor, createInteractor, router)
	return userController
}

// Injector.mock.go:

var controllerSet = wire.NewSet(v1.NewUserController)

var repositorySet = wire.NewSet(inmemory.NewUserRepository, wire.Bind(new(user3.Repository), new(*inmemory.UserRepository)))

var usecaseSet = wire.NewSet(user2.NewCreateInteractor, user.NewGetInteractor, wire.Bind(new(user4.CreateUsecase), new(*user2.CreateInteractor)), wire.Bind(new(user4.GetUsecase), new(*user.GetInteractor)))
