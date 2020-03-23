//+build wireinject

package inject

import (
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"

	ap "worksample-go-layered-architecture/app/application/user"
	v1 "worksample-go-layered-architecture/app/controller/v1"
	user "worksample-go-layered-architecture/app/domain/user"
	pe "worksample-go-layered-architecture/app/persistence/inmemory"
	qs "worksample-go-layered-architecture/app/queryservice/inmemory/user"
	uc "worksample-go-layered-architecture/app/usecase/user"
)

// Setup is
func Setup(router *httprouter.Router) *v1.UserController {
	wire.Build(controllerSet, repositorySet, usecaseSet)

	return nil
}

var controllerSet = wire.NewSet(
	v1.NewUserController,
)

var repositorySet = wire.NewSet(
	pe.NewUserRepository,
	wire.Bind(new(user.Repository), new(*pe.UserRepository)),
)

var usecaseSet = wire.NewSet(
	ap.NewCreateInteractor,
	qs.NewGetInteractor,
	wire.Bind(new(uc.CreateUsecase), new(*ap.CreateInteractor)),
	wire.Bind(new(uc.GetUsecase), new(*qs.GetInteractor)),
)
