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
	fs       embed.FS
	shareDir string
)

func init() {
	flag.StringVar(&shareDir, "dir", "", "path to shared directory")
	flag.Parse()
	if len(shareDir) < 1 {
		log.Fatal("directory not set. -help?")
	}
}

func main() {
	if _, err := os.Stat(shareDir); os.IsNotExist(err) {
		log.Fatal(err)
	}

	if err := middleware.LoadWhitelist(); err != nil {
		log.Fatal(err)
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	templ := template.Must(template.New("").ParseFS(fs, "templates/admin/*"))
	router.SetHTMLTemplate(templ)
	router.Use(middleware.IPWhitelist())
	router.StaticFS("/file", http.Dir(shareDir))
	router.GET("/", handler.RedirectToFile)

	adminAuthorizedGroup := router.Group("/")
	adminAuthorizedGroup.Use(middleware.AdminList())
	adminAuthorizedGroup.GET("/admin", handler.AdminIndex)
	adminAuthorizedGroup.POST("/admin/:action", handler.AdminAction)

	router.Run(":8080")
}
