package controller

import (
	a "go-blog/bootstrap/lib/action"
	"log"
	"net/http"
)

type HomeController struct {
	endpoints map[string]a.IAction
}

func (controller *HomeController) InitEndpoints() {
	if controller.endpoints == nil {
		endpoints := map[string]a.IAction {
			"/": IndexAction,
			"/home": HomeAction,
			"/about-us": AboutUsAction,
		}
		controller.endpoints = endpoints
	}
}

func (controller *HomeController) GetEndpoints() map[string]a.IAction {
	controller.InitEndpoints()
	return controller.endpoints
}

var IndexAction = a.ProtoAction{UnderlyingRun: func(w http.ResponseWriter, r *http.Request) {
	log.Printf("Hello from my home action!")
}}

var HomeAction = a.ProtoAction{UnderlyingRun: func(w http.ResponseWriter, r *http.Request) {
	log.Printf("Hello from my home action!")
}}

var AboutUsAction = a.ProtoAction{UnderlyingRun: func(w http.ResponseWriter, r *http.Request) {
	log.Printf("Hello from my home action!")
}}

