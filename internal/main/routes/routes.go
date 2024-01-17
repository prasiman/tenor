package routes

import (
	"main/internal/main/controllers/maincontroller"
	"main/internal/main/controllers/usercontroller"

	"github.com/gin-gonic/gin"

	"main/internal/main/middlewares"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	MainRoutes(router)
	AuthRoutes(router)
	UserRoutes(router)

	router.Run(":3000")

	return router
}

func MainRoutes(router *gin.Engine) {
	router.GET("/", maincontroller.GetServiceDetail)
}

func AuthRoutes(router *gin.Engine) {
	authV1 := router.Group("/api/v1")
	{
		authV1.POST("/register", usercontroller.Register)
		authV1.POST("/auth", usercontroller.Login)
	}
}

func UserRoutes(router *gin.Engine) {
	userV1 := router.Group("/api/v1").Use(middlewares.UserAuthMiddleware())
	{
		userV1.GET("/me", usercontroller.GetProfile)
		userV1.GET("/limit", usercontroller.GetCreditLimit)
		userV1.GET("/contracts", usercontroller.GetAllContracts)
		userV1.GET("/contracts/:id", usercontroller.GetContractByID)
	}
}
