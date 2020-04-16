package bootstrap

import (
	a "go-blog/bootstrap/lib/action"
	c "go-blog/bootstrap/lib/controller"
	"log"
)

type Registrar struct {}

func (registrar Registrar) RegisterControllerEndpoints(controllers ...c.IController) map[string]a.IAction {
	endpoints := map[string]a.IAction{}

	for _, ctrl := range controllers {
		for k, v := range ctrl.GetEndpoints() {
			endpoints[k] = v
		}
	}

	log.Printf("Controllers registered successfully.\nAll registered endpoints: %v", endpoints)
	return endpoints
}
