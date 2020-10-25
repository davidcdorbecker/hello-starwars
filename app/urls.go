package app

import "hello-starwars/controllers"

func mapUrls() {
	router.GET("/challenge", controllers.SayHello)
	router.GET("/challenge/starwars", controllers.HelloStarWars)
}