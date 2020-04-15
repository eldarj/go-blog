package controller

import(
	a "go-blog/bootstrap/lib/action"
)

type IController interface {
	GetEndpoints() map[string]a.IAction
}
