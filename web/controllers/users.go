package controllers

import (
	"github.com/kataras/iris"
	"strconv"
)

type UsersController struct {
	Ctx iris.Context
}

/*func (c *UsersController) BeforeActivation(b mvc.BeforeActivation) {
}*/

func (c *UsersController) Get() string {
	return "Index of users"
}

func (c *UsersController) Post() string {
	return "Create a user"
}

func (c *UsersController) GetBy(id int64) string {
	return "User with ID: " + strconv.FormatInt(id, 10)
}
