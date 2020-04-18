package bootstrap

import (
	auth "go-blog/auth/google"
	c "go-blog/controller"
	ss "go-blog/server"
	"log"
)

func Bootstrap() {
	log.Println("Boostrapping static server...")
	staticServer := ss.StaticServer{}
	staticServer.ServeResources("/resources/","./www/resources/")
	staticServer.ServeStylesheets("/css/","./www/css/")
	staticServer.ServeStylesheets("/js/","./www/js/")

	log.Println("Bootstrapping controllers...")
	Router{}.Start(Registrar{}.RegisterControllerEndpoints(
		&c.HomeController{},
		&c.DashboardController{},
		&auth.AuthenticationService{},
	))
}
