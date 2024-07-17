package api

import (
	_ "API-Gateway/api/docs"
	"API-Gateway/api/handler"
	"API-Gateway/api/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @title User
// @version 1.0
// @description API Gateway
// @host localhost:8080
// BasePath: /
func Router(hand *handler.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	stories := router.Group("/api/v1/stories")
	stories.Use(middleware.AuthMiddleware)
	{
		stories.POST("", hand.CreateStory)
		stories.PUT("/:story_id", hand.UpdateStory)
		stories.DELETE("/:story_id", hand.DeleteStory)
		stories.GET("", hand.GetAllStories)
		stories.GET("/:story_id", hand.GetStory)
		stories.POST("/:story_id/comments", hand.CommentStory)
		stories.GET("/:story_id/comments", hand.GetCommentsOfStory)
		stories.POST("/:story_id/like", hand.Like)
	}
	itineraries := router.Group("/api/v1/itineraries")
	itineraries.Use(middleware.AuthMiddleware)
	{
		itineraries.POST("", hand.Itineraries)
		itineraries.PUT("/:itinerary_id", hand.UpdateItineraries)
		itineraries.DELETE("/:itinerary_id", hand.DeleteItineraries)
		itineraries.GET("", hand.GetItineraries)
		itineraries.GET("/:itinerary_id", hand.GetItinerariesById)
		itineraries.POST("/:itinerary_id/comments", hand.CommentItineraries)

	}

	router.Use(middleware.AuthMiddleware)
	router.GET("/api/v1/destinations", hand.GetDestinations)
	router.GET("/api/v1/destinations/:destination_id", hand.GetDestinationsById)
	router.POST("/api/v1/messages", hand.SendMessage)
	router.GET("/api/v1/messages", hand.GetMessages)
	router.POST("/api/v1/travel-tips", hand.CreateTips)
	router.GET("/api/v1/travel-tips", hand.GetTips)
	router.GET("/api/v1/users/:user_id/statistics", hand.GetUserStat)
	router.GET("/api/v1/trending-destinations", hand.TopDestinations)

	return router
}
