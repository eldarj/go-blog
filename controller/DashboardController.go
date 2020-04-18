package controller

import (
	a "go-blog/bootstrap/lib/action"
	v "go-blog/bootstrap/lib/view"
)

type DashboardController struct{}

func (controller *DashboardController) GetEndpoints() map[string]a.IAction {
	return map[string]a.IAction{
		"/dashboard": dashboardAction,
	}
}

var dashboardAction = a.ViewAction{UnderlyingRun: func() string {
	return v.View("dashboard")
}}
