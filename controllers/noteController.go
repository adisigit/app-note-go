package controllers

import (
	"app-note-go/dto"
	"app-note-go/initializer"
	"app-note-go/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Tags Notes
// @Param body body dto.NoteCreateRequest true "Create note payload"
// @Router /notes/create [post]
// @Security ApiKeyAuth
func CreateNote(c *gin.Context) {
	var body dto.NoteCreateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	userID := c.MustGet("userID").(string)
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	note := models.Note{Title: body.Title, Content: body.Content, UserID: parsedUserID}
	result := initializer.DB.Create(&note)
	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.AbortWithStatusJSON(200, gin.H{"message": "Note created successfully"})
}

// @Tags Notes
// @Param page query int true "Page number"
// @Router /notes/pagination [get]
// @Security ApiKeyAuth
func GetNotePagination(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	limit := 10
	offset := (page - 1) * limit
	userID := c.MustGet("userID").(string)
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	var notes []struct {
		ID      string `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	var total int64
	initializer.DB.Model(&models.Note{}).Where("user_id = ?", parsedUserID).Count(&total)

	result := initializer.DB.Table("notes").Select("id", "title", "content").Where("user_id = ? AND deleted_at is null", parsedUserID).Limit(limit).Offset(offset).Find(&notes)
	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.AbortWithStatusJSON(200, gin.H{"data": gin.H{
		"notes":   notes,
		"page":    page,
		"limit":   limit,
		"total":   total,
		"hasNext": int64(offset+limit) < total,
	}})
}

// @Tags Notes
// @Param id path string true "Note ID"
// @Router /notes/{id} [get]
// @Security ApiKeyAuth
func GetNote(c *gin.Context) {
	id := c.Params.ByName("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	note := models.Note{}
	result := initializer.DB.Where("id = ?", parsedID).First(&note)
	if result.Error != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "Note not found"})
		return
	}
	c.AbortWithStatusJSON(200, gin.H{"data": note})
}

// @Tags Notes
// @Param body body dto.NoteUpdateRequest true "Update note payload"
// @Router /notes/update [put]
// @Security ApiKeyAuth
func UpdateNote(c *gin.Context) {
	var body dto.NoteUpdateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	note := models.Note{}
	result := initializer.DB.Where("id = ? AND user_id = ?", body.ID, c.MustGet("userID").(string)).First(&note)
	if result.Error != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "Note not found"})
		return
	}
	note.Title = body.Title
	note.Content = body.Content
	result = initializer.DB.Save(&note)
	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.AbortWithStatusJSON(200, gin.H{"message": "Note updated successfully"})
}

// @Tags Notes
// @Param id path string true "Note ID"
// @Router /notes/delete/{id} [delete]
// @Security ApiKeyAuth
func DeleteNote(c *gin.Context) {
	id := c.Params.ByName("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	note := models.Note{}
	result := initializer.DB.Where("id = ? AND user_id = ?", parsedID, c.MustGet("userID").(string)).Delete(&note)
	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(404, gin.H{"error": "Note not found"})
		return
	}
	c.AbortWithStatusJSON(200, gin.H{"message": "Note deleted successfully"})
}
