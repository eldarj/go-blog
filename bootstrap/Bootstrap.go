package bootstrap

import (
	c "go-blog/controller"
	"log"
)

func Bootstrap() {
	log.Println("Bootstrapping...")
	Router{}.Start(Registrar{}.RegisterControllerEndpoints(
		&c.HomeController{},
	))
}