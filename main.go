package main

import (
	"github.com/gin-gonic/gin"
	"gologin/handler"
	"gologin/internal"
	"gologin/log"
	"net/http"
)

func init() {
	internal.InitDB()
	log.InitLogger()
}

func initStaticRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "login for yourself",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "login for yourself",
		})
	})
	r.GET("/registry", func(c *gin.Context) {
		c.HTML(http.StatusOK, "registry.html", gin.H{
			"title": "registry your account",
		})
	})
	return r
}

func main() {
	r := initStaticRouter()

	r.GET("/activation", handler.ActivationAccount)
	r.POST("/create", handler.CreateAccount)
	r.POST("/login", handler.LoginAccount)

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
