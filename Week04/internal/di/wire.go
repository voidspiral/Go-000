// +build wireinject
// The build tag makes sure the stub is not built in the final build

package di

import (
	"school/internal/biz"
	"school/internal/data"

	service "school/internal/service"

	"github.com/google/wire"
)

/*
func InitApp() *App {
	panic(wire.Build(data.NewUser, biz.NewUser, service.NewSchool, NewApp))
}
*/

var Set = wire.NewSet(
	data.NewUser,
	wire.Bind(new(biz.UserPO), new(*data.User)),
	biz.NewUser,
	wire.Bind(new(service.UserBiz), new(*biz.User)),
	service.NewSchool,
)

func InitApp() *App {
	panic(wire.Build(Set, NewApp))
}
