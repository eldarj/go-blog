package controller

import(
	a "go-blog/bootstrap/lib/action"
)

type IController interface {
	InitEndpoints()
	GetEndpoints() map[string]*a.IAction
}
