package main

import (
	"embed"
	"fileShareByIP/handler"
	"fileShareByIP/middleware"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	//go:embed templates/*
	fs          embed.FS
	shareDir    string
	allowUpload bool
	port        string
)

func init() {
	flag.StringVar(&shareDir, "dir", "", "path to shared directory")
	flag.BoolVar(&allowUpload, "upl", false, "allow users to upload file")
	flag.StringVar(&port, "p", "8080", "port to listen")
	flag.Parse()
	if len(shareDir) < 1 {
		log.Fatal("share directory not set. try -help?")
	}
}

func main() {
	if _, err := os.Stat(shareDir); os.IsNotExist(err) {
		log.Fatalf("share directory is not exist ERR: %v", err)
	}

	if err := middleware.LoadWhitelist(); err != nil {
		log.Fatal(err)
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(func(c *gin.Context) {
		c.Set("shareDir", shareDir)
		c.Next()
	})
	router.Use(middleware.IPWhitelist())
	router.StaticFS("/file", http.Dir(shareDir))

	templ := template.Must(template.New("").ParseFS(fs, "templates/*"))
	router.SetHTMLTemplate(templ)

	router.GET("/", handler.RedirectToFile)
	if allowUpload {
		router.GET("/upload", handler.FileUploadPage)
		router.POST("/upload", handler.UploadFile)
	}
	adminAuthorizedGroup := router.Group("/")
	adminAuthorizedGroup.Use(middleware.AdminList())
	adminAuthorizedGroup.GET("/admin", handler.AdminIndex)
	adminAuthorizedGroup.POST("/admin/:action", handler.AdminAction)

	router.Run(":" + port)
}
