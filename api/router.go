// package api

// import (
//     "github.com/gin-gonic/gin"
//     "Auth-Service/handlers"
//     "Auth-Service/middleware"
// )

// // SetupRouter configures the Gin router with routes and middleware
// func SetupRouter() *gin.Engine {
//     // Initialize Gin router
//     r := gin.Default()

//     // Apply middleware for JWT authentication
//     r.Use(middleware.AuthMiddleware())

//     // Routes that do not require authentication
//     r.POST("/login", handlers.Login)
//     r.POST("/verify-email", handlers.VerifyEmail)

//     // Routes that require authentication
//     auth := r.Group("/")
//     {
//         auth.POST("/register", handlers.Register)
//         auth.POST("/profile", handlers.Profile)
//         auth.POST("/update-profile", handlers.UpdateProfile)
//         auth.POST("/get-users", handlers.GetUsers)
//         auth.POST("/delete-user", handlers.DeleteUser)
//         auth.POST("/reset-password", handlers.ResetPassword)
//         auth.POST("/refresh-token", handlers.RefreshToken)
//         auth.POST("/logout", handlers.Logout)
//         auth.POST("/followers", handlers.GetFollowersByUserID)
//     }

//     return r
// }

package api

import (
	handlers "API-Gateway/api/handler"
	a "API-Gateway/genproto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// @title API Service
// @version 1.0
// @description API service
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(conn *grpc.ClientConn) *gin.Engine {
	r := gin.Default()
	contentService := a.NewContentServiceClient(conn)

	st := handlers.Handler{ContentService: contentService}

	auth := r.Group("/content")
	{
		auth.POST("/createStory", st.CreateStory)
		auth.DELETE("/deleteStory", st.DeleteStory)
		auth.PUT("/updateStory", st.UpdateStory)
		auth.DELETE("/get_all", st.GetAllStories)
		// auth.POST("/reset-password", st.ResetPassword)
		// auth.POST("/refresh-token", st.RefreshToken)
		// auth.POST("/logout", st.Logout)
		// auth.GET("/followers", st.GetFollowersByUserID)
	}

	return r
}
