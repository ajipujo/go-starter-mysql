package routes

import (
	"gostart/auth"
	"gostart/handler"
	"gostart/helper"
	"gostart/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(route *gin.Engine) {
	authService := auth.NewService()
	userService := user.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	api := route.Group("/api/v1")

	api.POST("/", Welcome)
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/checkemailavailability", userHandler.CheckEmailAvailability)
	api.POST("/avatars", auth.AuthMiddleware(authService, userService), userHandler.UploadAvatar)
	api.PUT("/users", auth.AuthMiddleware(authService, userService), userHandler.UpdateUser)
}

func Welcome(c *gin.Context) {
	response := helper.ApiResponse("Welcome to Golang Starter", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
