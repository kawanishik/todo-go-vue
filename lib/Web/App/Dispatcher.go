package web_app_dispatcher

import (
    // "fmt"
	// "net/http"
	"github.com/kawanishik/todo-go-vue/lib/Service"
    "github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
        c.Next()
    }
}

func HandleRoutes() {
    dispatcher := gin.Default()
    dispatcher.Use(CORSMiddleware()) // CORSミドルウェアを適用
    dispatcher.GET("/api/index", service_todo.Index)
    dispatcher.POST("/api/create", service_todo.Create)
    dispatcher.DELETE("/api/delete", service_todo.Delete)
    dispatcher.PUT("/api/edit", service_todo.Edit)
    dispatcher.Run(":3000")
}