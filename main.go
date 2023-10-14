package main

import (


	"tugas8/controllers"
	"tugas8/connect"
	"github.com/labstack/echo/v4"
)

func main(){
	config.InitDatabase()
	e := echo.New()
e.POST("/register", controller.RegisterController)


	e.POST("/pekerjaans", controller.CreatePekerjaansController)
	e.GET("/pekerjaans", controller.GetPekerjaansController)
	e.DELETE("/pekerjaans/:id", controller.DeletePekerjaans)
	e.PUT("/pekerjaans/:id", controller.UpdatePekerjaans)


	e.POST("/warga", controller.CreateWargasController)
	e.GET("/warga", controller.GetWargasController)
	e.Start(":8080")
}
