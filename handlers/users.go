package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/primekobie/lucy/models"
	"github.com/primekobie/lucy/services"
)

var ErrServerError = errors.New("server failed to process your request")

func (h *ServiceHandler) CreateUser(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8,max=20"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := h.users.CreateUser(c.Request.Context(), input.Name, input.Email, input.Password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateUser) {
			c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (h *ServiceHandler) GetUser(c *gin.Context) {
	userId := c.Param("id")
	if err := validate.Var(userId, "uuid"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	user, err := h.users.FetchUser(c.Request.Context(), uuid.MustParse(userId))
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *ServiceHandler) UpdateUserData(c *gin.Context) {
	var input map[string]any
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	idString, ok := c.Get("user_id")
	if !ok {
		slog.Error("failed to fetch user id from context")
		c.JSON(http.StatusInternalServerError, gin.H{"message": ErrServerError.Error()})
		return
	}

	id := uuid.MustParse(idString.(string))
	input["id"] = id

	user, err := h.users.UpdateUser(c.Request.Context(), input)
	if err != nil {
		if errors.Is(err, services.ErrFailedOperation) {
			c.JSON(http.StatusInternalServerError, gin.H{"message": ErrServerError.Error()})
			return
		}
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}


func (h *ServiceHandler) DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	if err := validate.Var(userId, "uuid"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	err := h.users.DeleteUser(c.Request.Context(), uuid.MustParse(userId))
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
