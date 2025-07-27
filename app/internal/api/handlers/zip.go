package handlers

import (
	"net/http"
	"path"
	"zip-app/internal/schemas"
	"zip-app/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ZipHandler struct {
	service service.ZipService
}

func NewHandler(serv service.ZipService) ZipHandler {
	return ZipHandler{
		service: serv,
	}
}

func (h *ZipHandler) CreateTask(c *gin.Context) {
	ip, userAgent := c.ClientIP(), c.GetHeader("User-Agent")
	
	res, err := h.service.CreateTask(ip, userAgent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": res})
}

func (h *ZipHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")

	idParse, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id input format(must be UUID)"})
		return
	}

	var filesToAdd schemas.InsertFiles

	if err = c.ShouldBindJSON(&filesToAdd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input files"})
		return
	}

	for _, file := range filesToAdd.Links {
		if path.Ext(file) != ".jpg" || path.Ext(file) != ".pdf" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "only .pdf or .jpg files"})
			return
		}
	}

	err = h.service.UpdateTask(idParse, filesToAdd.Links)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (h *ZipHandler) CheckStatus(c *gin.Context) {
	id := c.Param("id")

	idParse, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id input format(must be UUID)"})
		return
	} 
	res, err := h.service.CheckStatus(idParse)

	c.JSON(http.StatusOK, gin.H{"message": res})
}