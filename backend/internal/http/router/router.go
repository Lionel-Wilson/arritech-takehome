package router

import (
	"net/http"

	"github.com/Lionel-Wilson/arritech-takehome/internal/api/user"
	internaluser "github.com/Lionel-Wilson/arritech-takehome/internal/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func New(
	logger *logrus.Logger,
	userService internaluser.Service,
) http.Handler {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	userHandler := user.NewUserHandler(logger, userService)

	{
		v1 := router.Group("api/v1")
		{
			users := v1.Group("/users")
			{
				users.GET("/", userHandler.GetUsers())
				users.GET("/:id", userHandler.GetUser())
				users.POST("/", userHandler.CreateUser())
				users.PATCH("/:id", userHandler.UpdateUser())
				users.DELETE("/:id", userHandler.DeleteUser())
			}
		}
	}

	return router
}
