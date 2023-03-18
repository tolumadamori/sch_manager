package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/tolumadamori/sch_manager/docs"
	"github.com/tolumadamori/sch_manager/pkg/routes"
)

// @title 	School Management Server
// @version	1.0
// @description A School Management Server

// @host 	localhost:8001
// @BasePath /
func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.CourseRoutes(r)
	routes.StudentRoutes(r)
	routes.AdminRoutes(r)
	log.Fatal(r.Run(":8080"))
}
