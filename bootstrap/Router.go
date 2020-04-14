package bootstrap

import (
	i "go-blog/bootstrap/interface"
	"net/http"
)

type Router struct {
}

func (router Router) Start(registrar Registrar) {
	for uri, endpoint := range registrar.GetEndpoints() {
		router.listen(uri, endpoint)
	}
}

func (router Router) listen(uri string, action i.IAction) {
	http.HandleFunc(uri, action.Do)
}