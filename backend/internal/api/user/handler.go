package user

import (
	"errors"
	"github.com/Lionel-Wilson/arritech-takehome/internal/user/storage"
	"net/http"
	"strconv"

	"github.com/Lionel-Wilson/arritech-takehome/internal/api/user/dto"
	dtomapper "github.com/Lionel-Wilson/arritech-takehome/internal/api/user/dto/mapper"
	"github.com/Lionel-Wilson/arritech-takehome/internal/pkg/http/mapper"
	"github.com/Lionel-Wilson/arritech-takehome/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler interface {
	GetUsers() gin.HandlerFunc
	GetUser() gin.HandlerFunc
	UpdateUser() gin.HandlerFunc
	DeleteUser() gin.HandlerFunc
	CreateUser() gin.HandlerFunc
}

type handler struct {
	logger      *logrus.Logger
	userService user.Service
}

func NewUserHandler(
	logger *logrus.Logger,
	userService user.Service,
) Handler {
	return &handler{
		logger:      logger,
		userService: userService,
	}
}

func (h *handler) GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (h *handler) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			h.logger.WithContext(ctx).Errorf("Invalid user ID: %v", err)
			c.JSON(http.StatusBadRequest, mapper.ToErrorResponse("invalid user ID"))
		}

		userDomain, err := h.userService.GetUser(ctx, userID)
		if err != nil {
			h.logger.WithContext(ctx).Errorf("Failed to get user: %v", err)
			c.JSON(MapErrorToStatusCodeAndMessage(err))

			return
		}

		c.JSON(http.StatusOK, dtomapper.MapUserToResponse(userDomain))
	}
}

func (h *handler) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			h.logger.WithContext(ctx).Errorf("Invalid user ID: %v", err)
			c.JSON(http.StatusBadRequest, mapper.ToErrorResponse("invalid user ID"))
		}

		var updateUserRequest dto.UpdateUserRequest
		err = c.ShouldBindJSON(&updateUserRequest)
		if err != nil {
			h.logger.WithContext(ctx).Errorf("Invalid request: %v", err)
			c.JSON(http.StatusBadRequest, mapper.ToErrorResponse("invalid request"))

			return
		}

		userObject := dtomapper.MapUpdateUserRequestToDomain(updateUserRequest, userID)

		err = h.userService.UpdateUserDetails(ctx, userObject)
		if err != nil {
			h.logger.WithContext(ctx).Errorf("Failed to update user details: %v", err)
			c.JSON(MapErrorToStatusCodeAndMessage(err))

			return
		}

		c.JSON(http.StatusOK, mapper.ToSimpleMessageResponse("user details updated"))
	}
}

func (h *handler) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			h.logger.WithContext(ctx).Errorf("Invalid user ID: %v", err)
			c.JSON(http.StatusBadRequest, mapper.ToErrorResponse("invalid user ID"))
		}

		err = h.userService.DeleteUser(ctx, userID)
		if err != nil {
			h.logger.WithContext(ctx).Errorf("Failed to delete user: %v", err)
			c.JSON(MapErrorToStatusCodeAndMessage(err))

			return
		}

		c.JSON(http.StatusOK, mapper.ToSimpleMessageResponse("user deleted"))
	}
}

func (h *handler) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		var createUserRequest dto.CreateUserRequest

		err := c.ShouldBindJSON(&createUserRequest)
		if err != nil {
			h.logger.WithContext(ctx).Errorf("Invalid request: %v", err)
			c.JSON(http.StatusBadRequest, mapper.ToErrorResponse("invalid request"))

			return
		}

		userObject := dtomapper.MapCreateUserRequestToDomain(createUserRequest)

		err = h.userService.CreateUser(ctx, userObject)
		if err != nil {
			h.logger.WithContext(ctx).Errorf("Failed to create user: %v", err)
			c.JSON(MapErrorToStatusCodeAndMessage(err))

			return
		}

		c.JSON(http.StatusOK, mapper.ToSimpleMessageResponse("user created"))
	}
}

func MapErrorToStatusCodeAndMessage(err error) (code int, obj any) {
	switch {
	case errors.Is(err, storage.ErrUserNotFound):
		return http.StatusNotFound, mapper.ToErrorResponse("User not found")
	case errors.Is(err, storage.ErrUserEmailExists):
		return http.StatusBadRequest, mapper.ToErrorResponse("User email already exists")
	case errors.Is(err, user.ErrUserMustBeAtLeast18YearsOld):
		return http.StatusBadRequest, mapper.ToErrorResponse("User must be at least 18 years old")
	default:
		return http.StatusInternalServerError, mapper.ToErrorResponse("Something went wrong")
	}
}
