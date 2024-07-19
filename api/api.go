package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/mirjalilova/api_gateway/api/handlers"
	"github.com/mirjalilova/api_gateway/api/middleware"
	// middleware "github.com/mirjalilova/api_gateway/api/middleware"
)

// @title Time Capsule API
// @version 1.0
// @description API for Time Capsule resources
// @host localhost:8080
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func Engine(handler *handlers.Handlers) *gin.Engine {
	router := gin.Default()

	router.GET("api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust for your specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// protected := router.Group("/", middleware.JWTMiddleware())

	router.Use(middleware.JWTMiddleware())
	user := router.Group("/user")
	{
		user.GET("/profiles", handler.GetProfile)
        user.PUT("/profiles", handler.EditProfile)
        user.PUT("/passwords", handler.ChangePassword)
        user.GET("/setting", handler.GetSetting)
		user.PUT("/setting", handler.EditSetting)
		user.DELETE("/", handler.DeleteUser)
	}

	memory := router.Group("/memories")
	{
		memory.GET("/all", handler.GetAllMemories)
		memory.GET("/", handler.GetMemory)
		memory.POST("", handler.CreateMemory)
		memory.PUT("/", handler.UpdateMemory)
		memory.DELETE("/", handler.DeleteMemory)
		memory.GET("/hictorical", handler.GetHistoricalMemory)
		memory.GET("/tag", handler.GetByTagMemory)
		memory.GET("/others", handler.GetMemoriesOfOthers)
	}

	media := router.Group("/media")
	{
		media.GET("/", handler.GetMedia)
        media.POST("/", handler.CreateMedia)
        media.DELETE("/", handler.DeleteMedia)
	}

	comment := router.Group("/comments")
	{
		comment.GET("/", handler.GetComments)
        comment.POST("/", handler.CreateComment)
        comment.PUT("/", handler.UpdateComment)
        comment.DELETE("/", handler.DeleteComment)
	}

	share := router.Group("/shares")
	{
        share.GET("/", handler.GetShares)
        share.PUT("/", handler.Updateshare)
		share.POST("/", handler.CreateShare)
    }

	millistone := router.Group("/millistones")
	{
		millistone.POST("/", handler.CreateMillistones)
		millistone.PUT("/:id", handler.UpdateMillistones)
		millistone.DELETE("/:id", handler.DeleteMillistones)
		millistone.GET("/:id", handler.GetMillistone)
		millistone.GET("/", handler.GetAllMillistones)
		millistone.GET("/date", handler.GetByDateMillistones)
	}

	return router
}
