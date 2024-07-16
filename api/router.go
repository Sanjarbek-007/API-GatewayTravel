package api

import (
	_ "API-Gateway/api/docs"
	handlers "API-Gateway/api/handler"
	"API-Gateway/api/middleware"
	genproto "API-Gateway/genproto"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"
	files "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// RouterApi @title API Service
// @version 1.0
// @description API service
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func RouterApi(con1 *grpc.ClientConn, con2 *grpc.ClientConn, logger *zap.Logger) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	paymentCon := genproto.NewContentServiceClient(con1)
	userCon := genproto.NewUserServiceClient(con2)
	h := handlers.NewHandler(paymentCon, userCon, logger)

	user := router.Group("/api/user")
	{
		user.GET("/profile", h.Profile)
		user.PUT("/update-profile", h.UpdateProfile)
	}

	authRoutes := router.Group("/")
	authRoutes.Use(middleware.AuthMiddleware())
	{
		content := authRoutes.Group("/content")
		{
			content.GET("/createStory", h.CreateStory)
			content.PUT("/updateStory/{id}", h.UpdateStory)
			content.DELETE("deleteStory/{id}", h.DeleteStory)
			content.GET("/getAllStories", h.GetAllStories)
			content.GET("storyFullInfo/{story_id}", h.StoryFullInfo)
			content.POST("/commentStory/{story_id}", h.CommentStory)
			content.GET("/getAllComments/{story_id}", h.GetAllStories)
			content.POST("createLike/{story_id}", h.CreateLike)

			content.POST("createItinerariesd", h.CreateItineraries)
			content.POST("/updateItineraries/{id}", h.UpdateItineraries)
			content.DELETE("/deleteItineraries/{id}", h.DeleteItineraries)
			content.GET("/getAllItineraries", h.GetAllItineraries)
			content.DELETE("itinerariesFullInfo/{itinerary_id}", h.ItinerariesFullInfo)
			content.POST("/commentItineraries/{itinerary_id}", h.CommentItineraries)
			content.GET("/getDestinations", h.GetDestinations)
			content.GET("/getDestinationInfo/{destination_id}", h.GetDestinationInfo)

			content.POST("sentMessage", h.SentMessage)
			content.GET("/getAllMessages", h.GetAllMessages)

			content.POST("/createTravelTip", h.CreateTravelTip)
			content.GET("/getTravelTips", h.GetTravelTips)
			content.GET("userStatistics/{user_id}", h.UserStatistics)
		}

		return router
	}
}
