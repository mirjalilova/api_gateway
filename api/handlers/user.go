package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	auth "github.com/mirjalilova/api_gateway/genproto/auth"
	_ "github.com/mirjalilova/api_gateway/api/docs"
)

// GetProfile godoc
// @Summary Get user profile
// @Description Retrieve the profile of a user with the specified ID
// @Tags profiles
// @Accept json
// @Produce json
// @Success 200 {object} auth.UserRes 
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /profiles [get]
func (h *Handlers) GetProfile(c *gin.Context) {
	req := &auth.GetById{
		Id: "567bc6be-1806-43a5-9d07-c386ef92b717",
	}

	profile, err := h.User.GetProfile(c, req)
	if err != nil {
		log.Println("Error getting profile:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

// EditProfile godoc
// @Summary Edit user profile
// @Description Update the profile of a user with the specified ID
// @Tags profiles
// @Accept json
// @Produce json
// @Param profile body auth.UserRes true "Updated profile details"
// @Success 200 {object} string "Profile updated successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /profiles [put]
func (h *Handlers) EditProfile(c *gin.Context) {
	var body auth.EditProfileReqBpdy
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	req := &auth.UserRes{
		Id:          "567bc6be-1806-43a5-9d07-c386ef92b717",
		Username:    body.Username,
		Email:       body.Email,
		FullName:    body.FullName,
		DateOfBirth: body.DateOfBirth,
	}

	input, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Producer.ProduceMessages("upd-user", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Profile for user %s updated successfully", req.Id)})
}

// ChangePassword godoc
// @Summary Change user password
// @Description Update the password of a user with the specified ID
// @Tags passwords
// @Accept json
// @Produce json
// @Param password body auth.ChangePasswordReqBody true "Updated password details"
// @Success 200 {object} string "Password updated successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /passwords [put]
func (h *Handlers) ChangePassword(c *gin.Context) {
	var body auth.ChangePasswordReqBody
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	req := &auth.ChangePasswordReq{
		Id:              "567bc6be-1806-43a5-9d07-c386ef92b717",
		CurrentPassword: body.CurrentPassword,
		NewPassword:     body.NewPassword,
	}

	input, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Producer.ProduceMessages("upd-pass", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Password for user %s updated successfully", req.Id)})
}

// GetSetting godoc
// @Summary Get user settings
// @Description Retrieve the settings of a user with the specified ID
// @Tags settings
// @Accept json
// @Produce json
// @Success 200 {object} auth.Setting
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /settings [get]
func (h *Handlers) GetSetting(c *gin.Context) {
	req := &auth.GetById{
		Id: "567bc6be-1806-43a5-9d07-c386ef92b717",
	}

	setting, err := h.User.GetSetting(c, req)
	if err != nil {
		log.Println("Error getting setting:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, setting)
}

// EditSetting godoc
// @Summary Edit user settings
// @Description Update the settings of a user with the specified ID
// @Tags settings
// @Accept json
// @Produce json
// @Param setting body auth.Setting true "Updated setting details"
// @Success 200 {object} string "Setting updated successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /settings [put]
func (h *Handlers) EditSetting(c *gin.Context) {
	var body auth.Setting
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	req := &auth.SettingReq{
		Id:           "567bc6be-1806-43a5-9d07-c386ef92b717",
		PrivacyLevel: body.PrivacyLevel,
		Notification: body.Notification,
		Language:     body.Language,
		Theme:        body.Theme,
	}

	input, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.Producer.ProduceMessages("upd-setting", input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Setting for user %s updated successfully", req.Id)})
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete a user with the specified ID
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} string "User deleted successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /users [delete]
func (h *Handlers) DeleteUser(c *gin.Context) {
	req := &auth.GetById{
		Id: "567bc6be-1806-43a5-9d07-c386ef92b717",
	}

	_, err := h.User.DeleteUser(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User %s deleted successfully", req.Id)})
}
