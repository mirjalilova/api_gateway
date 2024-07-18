package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	// md "github.com/mirjalilova/api_gateway/api/middleware"
	pb "github.com/mirjalilova/api_gateway/genproto/memory"

	"github.com/gin-gonic/gin"

	_ "github.com/mirjalilova/api_gateway/api/docs"
)

// @Summary Create a memory
// @Description Create a new memory
// @Tags memories
// @Accept json
// @Produce json
// @Param body body pb.MemoryCreateBody true "MemoryCreateBody"
// @Success 200 {object} string "Memory created successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /memories [post]
func (h *Handlers) CreateMemory(ctx *gin.Context) {
	var body pb.MemoryCreateBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// userID := getuserId(ctx)
	userID := "567bc6be-1806-43a5-9d07-c386ef92b717"

	req := &pb.MemoryCreate{
		UserId:      userID,
		Title:       body.Title,
		Description: body.Description,
		Date:        body.Date,
		Tags:        body.Tags,
		Locations:   body.Locations,
		PlaceName:   body.PlaceName,
		Type:        body.Type,
		Privacy:     body.Privacy,
	}

	input, err := json.Marshal(req)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	err = h.Producer.ProduceMessages("create-memory", input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Memory %s created successfully", req.Title)})
}

// @Summary Get a memory
// @Description Get a memory by ID
// @Tags memories
// @Accept json
// @Produce json
// @Param id query string true "Memory ID"
// @Success 200 {object} pb.MemoryRes
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /memories [get]
func (h *Handlers) GetMemory(ctx *gin.Context) {
	id := ctx.Query("id")

	req := &pb.GetById{
		Id:     id,
		// UserId: getuserId(ctx),
		UserId: "567bc6be-1806-43a5-9d07-c386ef92b717",
	}

	res, err := h.Memory.Get(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary Get all memories
// @Description Get all memories with optional filters
// @Tags memories
// @Accept json
// @Produce json
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Param tag query string false "Tag"
// @Param type query string false "Type"
// @Param user_id query string false "User ID"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} pb.GetAllRes
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error
// @Router /memories/all [get]
func (h *Handlers) GetAllMemories(ctx *gin.Context) {
	start_date := ctx.Query("start_date")
	end_date := ctx.Query("end_date")
	tag := ctx.Query("tag")
	typ := ctx.Query("type")
	user_id := ctx.Query("user_id")
	limitStr := ctx.Query("limit")
	offsetStr := ctx.Query("offset")

	var limit, offset int
	var err error
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
			return
		}
	} else {
		limit = 0
	}

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
			return
		}
	} else {
		offset = 0
	}

	if user_id == "" {
		// userID = getuserId(ctx)
		user_id = "567bc6be-1806-43a5-9d07-c386ef92b717"
	}

	req := &pb.GetAllReq{
		StartDate: start_date,
		EndDate:   end_date,
		Tag:       tag,
		Type:      typ,
		UserId:    user_id,
		Filter: &pb.Filter{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	}

	res, err := h.Memory.GetAll(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary Update a memory
// @Description Update a memory by ID
// @Tags memories
// @Accept json
// @Produce json
// @Param id path string true "Memory ID"
// @Param body body pb.MemoryCreateBody true "MemoryCreateBody"
// @Success 200 {object} string "Memory updated successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error
// @Router /memories [put]
func (h *Handlers) UpdateMemory(ctx *gin.Context) {
	id := ctx.Param("id")

	var body pb.MemoryCreateBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// userID := getuserId(ctx)
	userID := "567bc6be-1806-43a5-9d07-c386ef92b717"

	req := &pb.MemoryUpdate{
		Id:          id,
		UserId:      userID,
		Title:       body.Title,
		Description: body.Description,
		Date:        body.Date,
		Tags:        body.Tags,
		Locations:   body.Locations,
		PlaceName:   body.PlaceName,
		Type:        body.Type,
		Privacy:     body.Privacy,
	}

	input, err := json.Marshal(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.Producer.ProduceMessages("update-memory", input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Memory %s updated successfully", req.Title)})
}

// @Summary Delete a memory
// @Description Delete a memory by ID
// @Tags memories
// @Accept json
// @Produce json
// @Param id path string true "Memory ID"
// @Success 200 {object} string "Memory deleted successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error
// @Router /memories [delete]
func (h *Handlers) DeleteMemory(ctx *gin.Context) {
	id := ctx.Param("id")

	req := &pb.GetById{
		Id:     id,
		// UserId: getuserId(ctx),
		UserId: "567bc6be-1806-43a5-9d07-c386ef92b717",
	}

	_, err := h.Memory.Delete(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Memory %s deleted successfully", id)})
}

// func getuserId(ctx *gin.Context) string {
// 	user_id, err := md.GetUserId(ctx.Request)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return ""
// 	}
// 	return user_id
// }
