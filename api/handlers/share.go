package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mirjalilova/api_gateway/api/docs"
	pb "github.com/mirjalilova/api_gateway/genproto/memory"
)

// CreateShare godoc
// @Summary Create a share
// @Description Share a memory with other users
// @Tags shares
// @Accept json
// @Produce json
// @Param share body pb.ShareCreateBody true "Share Create Body"
// @Success 201 {object} string "Share created successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /shares [post]
func (h *Handlers) CreateShare(c *gin.Context) {
	var share pb.ShareCreateBody
	if err := c.ShouldBindJSON(&share); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newShare := &pb.ShareCreate{
		MemoryId:   share.MemoryId,
		UserId:     "567bc6be-1806-43a5-9d07-c386ef92b717",
		SharedWith: share.SharedWith,
	}

	_, err := h.Share.Share(c, newShare)
	if err != nil {
		log.Println("Error creating share:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newShare)
}

// GetShares godoc
// @Summary Get shares
// @Description Get all shares for a memory
// @Tags shares
// @Accept json
// @Produce json
// @Param memory_id query string true "Memory ID"
// @Success 200 {object} pb.ShareRes
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /shares [get]
func (h *Handlers) GetShares(c *gin.Context) {
	memoryId := c.Query("memory_id")

	req := &pb.GetById{
		Id:     memoryId,
		UserId: "567bc6be-1806-43a5-9d07-c386ef92b717",
	}

	shares, err := h.Share.Get(c, req)
	if err != nil {
		log.Println("Error getting shares:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shares)
}

// Unshare godoc
// @Summary Unshare a memory
// @Description Remove a share for a memory
// @Tags shares
// @Accept json
// @Produce json
// @Param share_id query string true "Share ID"
// @Param share body pb.ShareDeleteBody true "Share Delete Body"
/// @Success 200 {object} string "Share updated successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /shares [put]
func (h *Handlers) Updateshare(c *gin.Context) {
	shareId := c.Query("share_id")

	body := &pb.ShareDeleteBody{}
	if err := c.ShouldBindJSON(body); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req := &pb.ShareDelete{
		Id:          shareId,
		MemoryId:    body.MemoryId,
		UserId:      "567bc6be-1806-43a5-9d07-c386ef92b717",
		UnsharedWith: body.UnsharedWith,
	}

	_, err := h.Share.Updateshare(c, req)
	if err != nil {
		log.Println("Error unsharing:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Share unshared successfully"})
}
