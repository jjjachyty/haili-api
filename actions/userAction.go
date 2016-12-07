package actions

import (
	"errors"
	"haili/services"

	"net/http"

	"github.com/lunny/tango"
)

type UserAction struct {
	tango.Ctx
}

var userService services.UserService

func (a UserAction) Post() interface{} {

	username := a.Ctx.Form("username") // return string
	password := a.Ctx.Form("password") // return string

	users := userService.GetUsers(map[string]interface{}{"Name": username, "PassWord": password})
	if len(users) > 0 {
		return "登录成功"
	}
	return errors.New("something error")
}

func (a UserAction) Options() interface{} {
	return http.StatusOK
}
