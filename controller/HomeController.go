package controller

import (
	a "go-blog/bootstrap/lib/action"
	"log"
)

type HomeController struct {
	endpoints map[string]*a.IAction
}

func (controller *HomeController) InitEndpoints() {
	if controller.endpoints == nil {
		var act1 a.IAction = IndexAction()
		var act *a.IAction = &act1
		endpoints := map[string]*a.IAction {
			"/": act,
		}
		controller.endpoints = endpoints
	}
}

func (controller *HomeController) GetEndpoints() map[string]*a.IAction {
	controller.InitEndpoints()
	return controller.endpoints
}

func IndexAction() a.IAction {
	action := a.ProtoAction{}
	action.Set(func() {
		log.Println("Hello from my index action!")
	})
	return &action
}

func HomeAction() a.IAction {
	action := a.ProtoAction{}
	action.Set(func() {
		log.Printf("Hello from my home action! %v", action)
	})
	return &action
}

func AboutUsAction() a.IAction {
	action := a.ProtoAction{}
	action.Set(func() {
		log.Printf("Hello from my about-us action! %v", action)
	})
	return &action
}
