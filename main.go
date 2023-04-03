package main

import (
	"log"

	"example.com/projectCrud/controller"
	"example.com/projectCrud/initializers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error")
	}

	initializers.ConnectToDB()

}

func main() {
	r := gin.Default()
	r.POST("/servicios/proyectos/crear", controller.CreateProject)
	r.GET("/servicios/proyectos", controller.FindAllProjects)
	r.GET("/servicios/proyectos/:id", controller.FindProjectById)
	r.GET("/servicios/proyectos/titulo/:title", controller.FindProjectByTitle)
	r.GET("/servicios/proyectos/fecha/:date", controller.FindProjectByDate)
	r.GET("/servicios/proyectos/usuario/:user_id", controller.FindProjectByUser)
	r.GET("/servicios/proyectos/tipo/:type", controller.FindProjectByType)
	r.DELETE("/servicios/proyectos/eliminar", controller.DeleteAllProjects)
	r.DELETE("/servicios/proyectos/:id/eliminar", controller.DeleteProjectById)
	r.DELETE("/servicios/proyectos/titulo/:title/eliminar", controller.DeleteProjectByTitle)
	r.DELETE("/servicios/proyectos/fecha/:date/eliminar", controller.DeleteProjectByDate)
	r.DELETE("/servicios/proyectos/usuario/:user_id/eliminar", controller.DeleteProjectByUser)
	r.DELETE("/servicios/proyectos/tipo/:type/eliminar", controller.DeleteProjectByType)
	r.PUT("/servicios/proyectos/:id/editar", controller.UpdateProject)
	r.Run()
}
