package handlers

import (
    "fmt"
    "net/http"

    pb "github.com/mirjalilova/api_gateway/genproto/memory"
    "github.com/gin-gonic/gin"
    _ "github.com/mirjalilova/api_gateway/api/docs"
)

// CreateComment godoc
// @Summary Create a comment
// @Description Create a new comment on a memory
// @Tags comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param comment body pb.CommentCreateBody true "Comment Create Body"
// @Success 200 {object} string "Comment created successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /comments [post]
func (h *Handlers) CreateComment(ctx *gin.Context) {
	var body pb.CommentCreateBody
    if err := ctx.ShouldBindJSON(&body); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := "567bc6be-1806-43a5-9d07-c386ef92b717"

    req := &pb.CommentCreate{
        UserId:   userID,
        MemoryId: body.MemoryId,
        Content:  body.Content,
    }

    _, err := h.Comment.Create(ctx, req)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Comment created successfully"})
}

// GetComments godoc
// @Summary Get comments
// @Description Get all comments for a memory
// @Tags comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param memory_id query string true "Memory ID"
// @Success 200 {object} pb.Comment
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /comments [get]
func (h *Handlers) GetComments(ctx *gin.Context) {
	memoryId := ctx.Query("memory_id")

    req := &pb.GetById{
        Id:     memoryId,
        UserId: "567bc6be-1806-43a5-9d07-c386ef92b717",
    }

    res, err := h.Comment.Get(ctx, req)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, res)
}

// UpdateComment godoc
// @Summary Update a comment
// @Description Update the content of an existing comment
// @Tags comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query string true "Comment ID"
// @Param comment body pb.CommentUpdateBody true "Comment Update Body"
// @Success 200 {object} string "Comment updated successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /comments [put]
func (h *Handlers) UpdateComment(ctx *gin.Context) {
	id := ctx.Query("id")
    var body pb.CommentUpdateBody
    if err := ctx.ShouldBindJSON(&body); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	req := &pb.CommentUpdate{
		Id:      id,
        Content: body.Content,
	}
    _, err := h.Comment.Update(ctx, req)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}

// DeleteComment godoc
// @Summary Delete a comment
// @Description Delete a comment by ID
// @Tags comments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query string true "Comment ID"
// @Success 200 {object} string "Comment deleted successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /comments [delete]
func (h *Handlers) DeleteComment(ctx *gin.Context) {
	id := ctx.Query("id")

    req := &pb.GetById{
        Id:     id,
        UserId: "567bc6be-1806-43a5-9d07-c386ef92b717",
    }

    _, err := h.Comment.Delete(ctx, req)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Comment %s deleted successfully", id)})
}
