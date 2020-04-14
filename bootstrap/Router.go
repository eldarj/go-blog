package bootstrap

import (
	a "go-blog/bootstrap/lib/action"
	"net/http"
)

type Router struct {
}

func (router Router) Start(registrar *Registrar) {
	for uri, endpoint := range registrar.GetEndpoints() {
		router.listen(uri, endpoint)
	}
	_ = http.ListenAndServe(":8080", nil)
}

func (router Router) listen(uri string, action a.IAction) {
	http.HandleFunc(uri, action.Run)
}