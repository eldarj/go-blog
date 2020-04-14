package bootstrap

import (
	a "go-blog/bootstrap/lib/action"
	c "go-blog/bootstrap/lib/controller"
	"log"
)

type Registrar struct {
	endpoints map[string]*a.IAction
}

func (registrar Registrar) Register(controllers ...c.IController) *Registrar {
	registrar.endpoints = map[string]*a.IAction{}

	for _, ctrl := range controllers {
		ctrlEndpoints := ctrl.GetEndpoints()
		log.Printf("Registering controller: %+v", ctrl)

		for k, v := range ctrlEndpoints {
			log.Printf("Registering endpoint: %+v, %+v", k, v)
			registrar.endpoints[k] = v
		}
	}

	log.Printf("Controllers registered.\nAll registered endpoints: %v", registrar.endpoints)
	return &registrar
}

func (registrar Registrar) GetEndpoints() map[string]*a.IAction {
	return registrar.endpoints
}
