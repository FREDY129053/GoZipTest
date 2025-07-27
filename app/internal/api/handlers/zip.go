package handlers

import (
	"log"
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
	
	res := h.service.CreateTask(ip, userAgent)

	log.Println(res.Err)
	c.JSON(res.Code, res.Message)
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

	res := h.service.UpdateTask(idParse, filesToAdd.Links)

	log.Println(res.Err)
	c.JSON(res.Code, res.Message)
}

func (h *ZipHandler) CheckStatus(c *gin.Context) {
	id := c.Param("id")

	idParse, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id input format(must be UUID)"})
		return
	} 
	res := h.service.CheckStatus(idParse)

	log.Println(res.Err)
	c.JSON(res.Code, res.Message)
}