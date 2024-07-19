package handlers

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
    _ "github.com/mirjalilova/api_gateway/api/docs"
    pb "github.com/mirjalilova/api_gateway/genproto/memory"
)

// CreateMedia godoc
// @Summary Create media
// @Description Create a new media entry
// @Tags media
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param media body pb.MediaCreateBody true "Media Create Body"
// @Success 200 {object} string "Media created successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /media [post]
func (h *Handlers) CreateMedia(ctx *gin.Context) {
    var body pb.MediaCreateBody
    if err := ctx.ShouldBindJSON(&body); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := "567bc6be-1806-43a5-9d07-c386ef92b717"

    req := &pb.MediaCreate{
        UserId:   userID,
        Type:     body.Type,
        Url:      body.Url,
        MemoryId: body.MemoryId,
    }

    _, err := h.Media.Create(ctx, req)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Media created successfully"})
}

// GetMedia godoc
// @Summary Get media
// @Description Get a media entry by ID
// @Tags media
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query string true "Memory ID"
// @Success 200 {object} pb.Media
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /media [get]
func (h *Handlers) GetMedia(ctx *gin.Context) {
    id := ctx.Query("id")

    req := &pb.GetById{
        Id:     id,
        UserId: "567bc6be-1806-43a5-9d07-c386ef92b717",
    }

    res, err := h.Media.Get(ctx, req)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, res)
}

// DeleteMedia godoc
// @Summary Delete media
// @Description Delete a media entry by ID
// @Tags media
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id query string true "Media ID"
// @Success 200 {object} string "Media deleted successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /media [delete]
func (h *Handlers) DeleteMedia(ctx *gin.Context) {
    id := ctx.Query("id")

    req := &pb.GetById{
        Id:     id,
        UserId: "567bc6be-1806-43a5-9d07-c386ef92b717",
    }

    _, err := h.Media.Delete(ctx, req)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Media %s deleted successfully", id)})
}
