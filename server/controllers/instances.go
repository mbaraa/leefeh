package controllers

import (
	"salsa/controllers/apis"
	"salsa/controllers/sockets"
)

var bindables []Bindable = []Bindable{
	apis.NewExampleApi(),
	sockets.NewEchoSocket(),
}

func GetControllers() []Bindable {
	return bindables
}
