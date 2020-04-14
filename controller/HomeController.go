package controller

import (
	a "go-blog/bootstrap/lib/action"
	v "go-blog/bootstrap/lib/view"
	"log"
	"net/http"
)

type HomeController struct {
	endpoints map[string]a.IAction
}

func (controller *HomeController) InitEndpoints() {
	if controller.endpoints == nil {
		endpoints := map[string]a.IAction{
			"/":         indexAction,
			"/home":     homeAction,
			"/about-us": aboutUsAction,
		}
		controller.endpoints = endpoints
	}
}

func (controller *HomeController) GetEndpoints() map[string]a.IAction {
	controller.InitEndpoints()
	return controller.endpoints
}

var indexAction = a.ProtoAction{UnderlyingRun: func(w http.ResponseWriter, r *http.Request) {
	log.Printf("Hello from my home action!")
}}

var homeAction = a.ProtoAction{UnderlyingRun: func(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./www/home.html")
}}

var aboutUsAction = a.ViewAction{UnderlyingRun: func() string {
	return v.View("about-us")
}}
