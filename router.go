package main

import "haili/actions"
import "github.com/lunny/tango"

func Router(t *tango.Tango) {
	t.Any("/users/:name", new(actions.UserAction))
	t.Post("/register", new(actions.UserAction))

	t.Post("/smscode", new(actions.SMSVerCodeAction))
	t.Get("/smscode/phone/:phone/smscode/:smscode", new(actions.SMSVerCodeAction))
}
