package web_app_dispatcher

import (
    "time"
	"github.com/kawanishik/todo-go-vue/lib/Service"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func setCors(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

func HandleRoutes() {
    dispatcher := gin.Default()
    setCors(dispatcher) // CORSミドルウェアを適用
    dispatcher.GET("/api/index", service_todo.Index)
    dispatcher.POST("/api/create", service_todo.Create)
    dispatcher.DELETE("/api/delete", service_todo.Delete)
    dispatcher.PUT("/api/edit", service_todo.Edit)
    dispatcher.Run(":3000")
}