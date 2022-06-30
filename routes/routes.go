package routes

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vishal2457/redmine-server/controllers"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Print("in middleware")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// SetupRouter sets up the router.
func SetupRouter() *gin.Engine {
	r := gin.Default()
	// r.Use(CORSMiddleware())
	r.Use(cors.Default())

	v1 := r.Group("/api/v1/")

	users := v1.Group("/users")
	{
		users.GET("/", controllers.GetUsers)
		users.GET("/:id", controllers.GetUser)
		users.POST("/", controllers.CreateUser)
		users.PATCH("/:id", controllers.UpdateUser)
		users.DELETE("/:id", controllers.DeleteUser)
	}

	projects := v1.Group("/projects")
	{
		projects.GET("/", controllers.GetProjects)
	}

	task := v1.Group("/task")
	{
		task.POST("/", controllers.CreateTask)
		task.GET("/", controllers.GetTasks)
	}

	return r
}
