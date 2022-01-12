package main

import (
	// "skb/Controller"
	// "skb/Model"

	// "log"
	// "net/http"
	// "os"
	// "time"

	// "golang.org/x/crypto/bcrypt"

	// "github.com/gin-gonic/gin"

	// jwt "github.com/appleboy/gin-jwt/v2"
	"net/http"
	"os"
	"skb/Auth"
	"skb/Controller"

	"github.com/gin-gonic/gin"
)

// var identityKey = "id"

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT,DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	port := os.Getenv("PORT")
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())
	if port == "" {
		port = "8080"
	}
	r.StaticFS("/file", http.Dir("public"))
	r.POST("/upload", Controller.Upload)
	r.POST("/register", Controller.Register)
	r.GET(`/detail/:id`, Controller.DetailArtikel)
	r.GET(`/user`, Controller.GetUser)
	r.GET(`/paginate/:id`, Controller.Paginate)
	r.GET(`/get/all`, Controller.GetAllArtikel)
	r.POST("/create/artikel", Controller.CreateArtikel)
	r.PUT(`/edit/artikel/:id`, Controller.EditArtikel)
	r.DELETE(`/delete/artikel/:id`, Controller.DeleteArtikel)
	Auth.Authenticator()
}
