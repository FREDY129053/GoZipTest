package handlers

import (
	"log"
	"net/http"
	"os"
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

// CreateTask	godoc
// @Summary 	Create task
// @Description Create task for archiver
// @Tags		Zip-Archive
// @Produce		json
// @Success 	201 	{object} 	schemas.CreatedTask
// @Failure 	423 	{object}  	schemas.APIError
// @Router 		/zip_task	[post]
func (h *ZipHandler) CreateTask(c *gin.Context) {
	ip, userAgent := c.ClientIP(), c.GetHeader("User-Agent")

	res := h.service.CreateTask(ip, userAgent)

	log.Println(res.Err)
	c.JSON(res.Code, res.Message)
}

// UpdateTask	godoc
// @Summary 	Update task
// @Description Add files to task
// @Tags		Zip-Archive
// @Accept		json
// @Produce		json
// @Param		id		path	string	true	"Task ID"
// @Param		files	body	schemas.InsertFiles	true	"Files to add"
// @Success 	200 	{object} 	schemas.MessageAnswer
// @Failure 	400 	{object} 	schemas.APIError
// @Failure 	404 	{object} 	schemas.APIError
// @Failure 	422 	{object} 	schemas.APIError
// @Router 		/zip_task/{id}	[put]
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
		if path.Ext(file) != ".jpg" && path.Ext(file) != ".pdf" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "only .pdf or .jpg files"})
			return
		}
	}

	res := h.service.UpdateTask(idParse, filesToAdd.Links)

	log.Println(res.Err)
	c.JSON(res.Code, res.Message)
}

// CheckStatus	godoc
// @Summary 	Check task status
// @Description Check status of task if it has less than 3 files. If task has 3 files it returns **link** to archive
// @Tags		Zip-Archive
// @Produce		json
// @Param		id		path	string	true	"Task ID"
// @Success 	200 	{object} 	schemas.TaskStatus
// @Failure 	400 	{object} 	schemas.APIError
// @Failure 	404 	{object} 	schemas.APIError
// @Router 		/zip_task/{id}/status	[get]
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

// DownloadArchive	godoc
// @Summary 	Download archive
// @Description Download archive from link
// @Tags		Zip-Archive
// @Produce		json
// @Param		file		path	string	true	"Filename"
// @Success 	200 	{object} 	schemas.TaskStatus
// @Failure 	400 	{object} 	schemas.APIError
// @Failure 	404 	{object} 	schemas.APIError
// @Router 		/zip_task/download/{file}	[get]
func (h *ZipHandler) DownloadArchive(c *gin.Context) {
	filename := c.Param("file")
	filePath := "archives/" + filename
	log.Println(filePath)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/octet-stream")

	c.File(filePath)
}
