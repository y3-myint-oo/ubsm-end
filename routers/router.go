package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitConnection - Create Routers
func InitConnection() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost"},
	// 	AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))
	router.Use(cors.Default())
	v1 := router.Group("/api/v1/")
	{
		v1.POST("/login", login)

		//v1.POST("/", createTodo)
		//v1.GET("/", fetchAllTodo)
		//v1.GET("/:id", fetchSingleTodo)
		//v1.PUT("/:id", updateTodo)
		//v1.DELETE("/:id", deleteTodo)

		v1.POST("/unit", createUnit)
		v1.GET("/unit", fetchAllUnit)
		v1.GET("/unit:id", fetchSingleUnit)
		v1.PUT("/unit:id", updateUnit)
		v1.POST("/delete/unit", deleteUnit)
	}
	router.Run(":8081")
}
