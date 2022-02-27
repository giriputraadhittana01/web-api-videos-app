package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/giriadhittana01/golang-gin-poc/controller"
	"github.com/giriadhittana01/golang-gin-poc/middlewares"
	"github.com/giriadhittana01/golang-gin-poc/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setUpLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setUpLogOutput()

	// Default
	// server := gin.Default()

	// Modified
	server := gin.New()
	// server.Use(gin.Recovery(), gin.Logger())
	server.Use(
		gin.Recovery(),
		// middlewares.Logger(),
		gin.Logger(),
		middlewares.BasicAuth(),
		// gindump.Dump(),
	)

	// Load HTML
	server.LoadHTMLGlob("templates/*.html")

	// Contoh
	// server.GET("/api/test",func(ctx *gin.Context){
	// 	ctx.JSON(200, gin.H{
	// 		"message" : "Hello Giri Putra Adhittana",
	// 	})
	// })

	// Rest API

	// server.GET("/api/v1/videos", func(ctx *gin.Context) {
	// 	ctx.JSON(200, videoController.FindAll())
	// })
	// server.POST("/api/v1/videos", func(ctx *gin.Context) {
	// 	err := videoController.Save(ctx)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusOK, gin.H{
	// 			"message": err.Error(),
	// 		})
	// 	} else {
	// 		ctx.JSON(http.StatusOK, gin.H{
	// 			"message": "Success",
	// 		})
	// 	}
	// })

	apiRoutes := server.Group("/api/v1")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})
		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"message": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Success",
				})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	server.Run(":" + port)

}
