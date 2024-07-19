package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/mirjalilova/api_gateway/api/docs"
	pb "github.com/mirjalilova/api_gateway/genproto/timeline"
)

// CreateMillistones godoc
// @Summary Create a new millistone
// @Description Create a new millistone with the specified details
// @Tags millistones
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body pb.MillistonesCreateBody true "Millistoneobjecttobecreated"
// @Success 200 {object} string "Millistone created successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /millistones [post]
func (h *Handlers) CreateMillistones(c *gin.Context) {
	var body pb.MillistonesCreate
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	req := &pb.MillistonesCreate{
		Title:       body.Title,
		Date:        body.Date,
		Description: body.Description,
		Category:    body.Category,
		UserId:      "567bc6be-1806-43a5-9d07-c386ef92b717",
	}

	input, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Producer.ProduceMessages("create-millistones", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Millistone created successfully"})
}

// UpdateMillistones godoc
// @Summary Update an existing millistone
// @Description Update an existing millistone with the specified details
// @Tags millistones
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Millistone ID"
// @Param body body pb.MillistonesCreateBody true "Updated millistone details"
// @Success 200 {object} string "Millistone updated successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error
// @Router /millistones/{id} [put]
func (h *Handlers) UpdateMillistones(c *gin.Context) {
	id := c.Query("id")

	var body pb.MillistonesCreateBody
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	req := &pb.MillistonesUpdate{
		Id:          id,
		Title:       body.Title,
		Date:        body.Date,
		Description: body.Description,
		Category:    body.Category,
		UserId:      "567bc6be-1806-43a5-9d07-c386ef92b717",
	}

	input, err := json.Marshal(req)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	err = h.Producer.ProduceMessages("update-millistones", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Millistone updated successfully"})
}

// DeleteMillistones godoc
// @Summary Delete a millistone by ID
// @Description Delete a millistone by its ID
// @Tags millistones
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Millistone ID"
// @Success 200 {object} string "Millistone deleted successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error
// @Router /millistones/{id} [delete]
func (h *Handlers) DeleteMillistones(c *gin.Context) {
	id := c.Param("id")

	req := &pb.GetById{
		Id:     id,
		UserId: "567bc6be-1806-43a5-9d07-c386ef92b717",
	}

	_, err := h.Millistones.Delete(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Millistone deleted successfully"})
}

// GetMillistone godoc
// @Summary Get a millistone by ID
// @Description Retrieve details of a millistone by its ID
// @Tags millistones
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Millistone ID"
// @Success 200 {object} pb.Millistone
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error
// @Router /millistones/{id} [get]
func (h *Handlers) GetMillistone(c *gin.Context) {
	id := c.Param("id")

	req := &pb.GetById{
		Id:     id,
		UserId: "567bc6be-1806-43a5-9d07-c386ef92b717",
	}

	millistone, err := h.Millistones.Get(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, millistone)
}

// GetAllMillistones godoc
// @Summary Get all millistones based on filters
// @Description Retrieve a list of millistones based on optional filters like date, category, user ID, limit, and offset
// @Tags millistones
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param date query string false "Date filter (YYYY-MM-DD)"
// @Param category query string false "Category filter"
// @Param user_id query string false "User ID filter"
// @Param limit query int false "Limit number of results"
// @Param offset query int false "Offset number for pagination"
// @Success 200 {object} pb.GetAllRes
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error
// @Router /millistones [get]
func (h *Handlers) GetAllMillistones(c *gin.Context) {
	date := c.Query("date")
	category := c.Query("category")
	user_id := c.Query("user_id")
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	var limit, offset int
	var err error
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
			return
		}
	} else {
		limit = 0
	}

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
			return
		}
	} else {
		offset = 0
	}

	if user_id == "" {
		user_id = "567bc6be-1806-43a5-9d07-c386ef92b717"
	}

	req := &pb.GetAllReq{
		UserId:   user_id,
		Category: category,
		Date:     date,
		Filter: &pb.Filter{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	}

	millistones, err := h.Millistones.GetAll(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, millistones)
}

// GetByDateMillistones godoc
// @Summary Get millistones within a date range
// @Description Retrieve millistones within a specified date range for a given user ID
// @Tags millistones
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param from_date query string true "Start date of range (YYYY-MM-DD)"
// @Param to_date query string true "End date of range (YYYY-MM-DD)"
// @Param user_id query string false "User ID filter"
// @Success 200 {object} pb.GetAllRes
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /millistones/date [get]
func (h *Handlers) GetByDateMillistones(c *gin.Context) {
	toDate := c.Query("to_date")
	fromDate := c.Query("from_date")
	userId := c.Query("user_id")

	if fromDate == "" || toDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "from_date and to_date parameters are required"})
		return
	}

	if userId == "" {
		userId = "567bc6be-1806-43a5-9d07-c386ef92b717"
	}

	// Log the parameters for debugging
	fmt.Println("Parameters:", fromDate, toDate, userId)

	req := &pb.GetByDate{
		UserId:   userId,
		ToDate:   toDate,
		FromDate: fromDate,
	}

	millistones, err := h.Millistones.GetByDateMillistones(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, millistones)
}
