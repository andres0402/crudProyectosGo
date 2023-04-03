package controller

import (
	"net/http"
	"strings"

	"example.com/projectCrud/initializers"
	"example.com/projectCrud/model"
	"github.com/gin-gonic/gin"
)

func CreateProject(c *gin.Context) {

	var body struct {
		Title               string `form:"title" json:"title"`
		Project_description string `form:"project_description" json:"project_description"`
		Project_date        string `form:"project_date" json:"project_date"`
		Project_type        string `form:"project_type" json:"project_type"`
		User_id             int    `form:"user_id" json:"user_id"`
	}

	c.Bind(&body)
	projectType := strings.ToLower(body.Project_type)

	tipos := []string{"modificacion", "reparacion"}

	if !contains(tipos, projectType) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Ingrese un tipo de proyecto valido"})
		return
	}

	project := model.Project{Title: body.Title, Project_description: body.Project_description, Project_date: body.Project_date, Project_type: projectType, User_id: body.User_id}

	result := initializers.DB.Create(&project) // pass pointer of data to Create

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No se pudo crear el proyecto"})
		return
	}
	c.JSON(200, gin.H{
		"project": project,
	})
}

func FindAllProjects(c *gin.Context) {
	var projects []model.Project
	initializers.DB.Find(&projects)

	c.JSON(200, gin.H{
		"projects": projects,
	})

}

func FindProjectById(c *gin.Context) {
	id := c.Param("id")
	var project model.Project
	result := initializers.DB.First(&project, id)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No existe el proyecto"})
		return
	}

	c.JSON(200, gin.H{
		"project": project,
	})

}

func FindProjectByTitle(c *gin.Context) {
	title := c.Param("title")
	var projects []model.Project
	initializers.DB.Where("title = ?", title).Find(&projects)

	if len(projects) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No existen proyectos con este titulo"})
		return
	}

	c.JSON(200, gin.H{
		"projects": projects,
	})

}

func FindProjectByDate(c *gin.Context) {
	date := c.Param("date")
	var projects []model.Project
	initializers.DB.Where("project_date = ?", date).Find(&projects)

	if len(projects) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No existen proyectos para esa fecha"})
		return
	}

	c.JSON(200, gin.H{
		"projects": projects,
	})

}

func FindProjectByUser(c *gin.Context) {
	user := c.Param("user_id")
	var projects []model.Project
	initializers.DB.Where("user_id = ?", user).Find(&projects)

	if len(projects) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No existen proyectos para este usuario"})
		return
	}

	c.JSON(200, gin.H{
		"projects": projects,
	})

}

func FindProjectByType(c *gin.Context) {
	tipo := strings.ToLower(c.Param("type"))
	var projects []model.Project
	initializers.DB.Where("project_type = ?", tipo).Find(&projects)

	if len(projects) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No existen proyectos de este tipo"})
		return
	}

	c.JSON(200, gin.H{
		"projects": projects,
	})

}

func DeleteAllProjects(c *gin.Context) {

	results := initializers.DB.Exec("DELETE FROM projects")

	if results.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Se produjo un error al eliminar los proyectos"})
		return
	}

	c.JSON(200, gin.H{
		"exito": "Se han eliminado los proyectos",
	})

}

func DeleteProjectById(c *gin.Context) {

	id := c.Param("id")
	results := initializers.DB.Delete(&model.Project{}, id)

	if results.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Se produjo un error al eliminar el proyecto"})
		return
	}

	c.JSON(200, gin.H{
		"exito": "Se ha eliminado el proyecto",
	})

}

func DeleteProjectByTitle(c *gin.Context) {

	title := c.Param("title")
	results := initializers.DB.Where("title = ?", title).Delete(&model.Project{})

	if results.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Se produjo un error al eliminar los proyectos"})
		return
	}

	c.JSON(200, gin.H{
		"exito": "Se han eliminado los proyectos",
	})

}

func DeleteProjectByDate(c *gin.Context) {

	date := c.Param("date")
	results := initializers.DB.Where("project_date = ?", date).Delete(&model.Project{})

	if results.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Se produjo un error al eliminar los proyectos"})
		return
	}

	c.JSON(200, gin.H{
		"exito": "Se han eliminado los proyectos",
	})

}

func DeleteProjectByUser(c *gin.Context) {

	user := c.Param("user_id")
	results := initializers.DB.Where("user_id = ?", user).Delete(&model.Project{})

	if results.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Se produjo un error al eliminar los proyectos"})
		return
	}

	c.JSON(200, gin.H{
		"exito": "Se han eliminado los proyectos",
	})

}

func DeleteProjectByType(c *gin.Context) {

	tipo := c.Param("type")
	results := initializers.DB.Where("project_type = ?", tipo).Delete(&model.Project{})

	if results.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Se produjo un error al eliminar los proyectos"})
		return
	}

	c.JSON(200, gin.H{
		"exito": "Se han eliminado los proyectos",
	})

}

func UpdateProject(c *gin.Context) {

	var body struct {
		Title               string `form:"title" json:"title"`
		Project_description string `form:"project_description" json:"project_description"`
		Project_date        string `form:"project_date" json:"project_date"`
		Project_type        string `form:"project_type" json:"project_type"`
		User_id             int    `form:"user_id" json:"user_id"`
	}

	c.Bind(&body)
	projectType := strings.ToLower(body.Project_type)

	tipos := []string{"modificacion", "reparacion"}

	if !contains(tipos, projectType) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Ingrese un tipo de proyecto valido"})
		return
	}

	id := c.Param("id")
	var project model.Project
	initializers.DB.First(&project, id)
	results := initializers.DB.Model(&project).Updates(model.Project{Title: body.Title, Project_description: body.Project_description, Project_date: body.Project_date, Project_type: projectType, User_id: body.User_id})

	if results.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Se produjo un error al editar el proyecto"})
		return
	}

	c.JSON(200, gin.H{
		"project": project,
	})

}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
