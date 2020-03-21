//+build wireinject

package inject

import (
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"

	ap "worksample-go-layered-architecture/app/application/user"
	v1 "worksample-go-layered-architecture/app/controller/v1"
	"worksample-go-layered-architecture/app/domain/user"
	infra "worksample-go-layered-architecture/app/infrastructure/inmemory"
	uc "worksample-go-layered-architecture/app/usecase/user"
)

// Setup is
func Setup(router *httprouter.Router) *v1.UserController {
	wire.Build(
		v1.NewUserController,
		ap.NewGetInteractor,
		infra.NewUserRepository,
		wire.Bind(new(user.Repository), new(*infra.UserRepository)),
		wire.Bind(new(uc.GetUsecase), new(*ap.GetInteractor)),
	)
	return nil
}
