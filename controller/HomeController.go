package controller

import (
	i "go-blog/bootstrap/interface"
)

type HomeController struct {
	endpoints map[string]i.IAction
}

func (controller *HomeController) InitEndpoints() {
	if controller.endpoints == nil {
		controller.endpoints = map[string]i.IAction {
			"/": IndexAction(),
			"/home": HomeAction(),
			"/about-us": AboutUsAction(),
		}
	}
}

func (controller *HomeController) GetEndpoints() map[string]i.IAction {
	controller.InitEndpoints()
	return controller.endpoints
}

func IndexAction() i.IAction {
	return i.IAction{
		Filepath: "./www/index.html",
	}
}

func HomeAction() i.IAction {
	return i.IAction{
		Filepath: "./www/home.html",
	}
}

func AboutUsAction() i.IAction {
	return i.IAction{
		Filepath: "./www/about-us.html",
	}
}