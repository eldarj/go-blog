package bootstrap

import (
	a "go-blog/bootstrap/lib/action"
	"net/http"
)

type Router struct {}

func (router Router) Start(endpoints map[string]a.IAction) {
	for uri, endpoint := range endpoints {
		router.listen(uri, endpoint)
	}
	_ = http.ListenAndServe(":8080", nil)
}

func (router Router) listen(uri string, action a.IAction) {
	http.HandleFunc(uri, action.Run)
}