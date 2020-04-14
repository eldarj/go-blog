package bootstrap

import (
	i "go-blog/bootstrap/interface"
	"log"
)

type Registrar struct {
	endpoints map[string]i.IAction
}

func (registrar Registrar) Register(controllers ...i.IController) Registrar {
	registrar.endpoints = map[string]i.IAction{}

	for _, ctrl := range controllers {
		ctrlEndpoints := ctrl.GetEndpoints()
		log.Printf("Registering controller: %+v", ctrl)

		for k, v := range ctrlEndpoints {
			log.Printf("Registering endpoint: %+v, %+v", k, v)
			registrar.endpoints[k] = v
		}
	}

	log.Printf("Controllers registered.\nAll registered endpoints: %v", registrar.endpoints)
	return registrar
}

func (registrar Registrar) GetEndpoints() map[string]i.IAction {
	return registrar.endpoints
}
