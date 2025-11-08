package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *ServiceHandler) UploadFiles(c *gin.Context) {
	_, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

}
