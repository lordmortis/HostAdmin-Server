package controllers

import (
	context2 "context"
	"github.com/kataras/iris"
	"github.com/lordmortis/HostaAdmin-Server/datamodels_raw"
	"github.com/lordmortis/HostaAdmin-Server/services"
	"strconv"
)

type UsersController struct {
	Ctx iris.Context

	DBService services.DatabaseService
}

/*func (c *UsersController) BeforeActivation(b mvc.BeforeActivation) {
}*/

func (c *UsersController) Get() (results datamodels_raw.UserSlice) {
	slice, err := datamodels_raw.Users().All(context2.Background(), c.DBService.GetConnection())
	if err != nil {
		println(err.Error())
		return nil
	}

	println("woo")
	return slice
}

func (c *UsersController) Post() string {
	return "Create a user"
}

func (c *UsersController) GetBy(id int64) string {
	return "User with ID: " + strconv.FormatInt(id, 10)
}
