package v1

import (
	"auction/internal/model"
	"auction/internal/service"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initStaffRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth-staff")
	{
		auth.POST("/sign-in", h.staffSignIn)
		auth.POST("/refresh", h.staffSignIn)
	}
	staff := api.Group("/staff", h.staffIdentity)
	{
		staff.GET("/info", h.staffInfo)
		staff.POST("/info/:uuid", h.staffUpdateInfo)
	}
}

// Sign In
type staffSignInInput struct {
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"'`
}

type staffSignInResponse struct {
	AccessToken string `json:"accessToken"`
}

// Staff Info Update
type staffUpdateInfoInput struct {
	Name  string `json:"name" binding:"required,min=2,max=64"`
	Phone string `json:"phone" binding:"required,max=10"`
}

// @Summary Sign in
// @Tags Staff
// @Description Sign in to staff account
// @ModuleID staffSignIn
// @Accept json
// @Produce json
// @Param input body staffSignInInput true "Sign In"
// @Success 201 {string} string "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth-staff/sign-in [post]
func (h *Handler) staffSignIn(c *gin.Context) {
	var input staffSignInInput
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")

		return
	}

	res, err := h.services.Staff.SignIn(c.Request.Context(), service.StaffSignInInput{
		Email:    input.Email,
		Password: input.Password,
	})

	if err != nil {
		if errors.Is(err, model.ErrStaffNotFound) {
			newResponse(c, http.StatusBadRequest, err.Error())

			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.SetCookie("refresh_token", res.RefreshToken, 60*60, "/auth", "localhost", false, true)

	c.JSON(http.StatusOK, staffSignInResponse{
		AccessToken: res.AccessToken,
	})
}

// @Summary Staff Info
// @Tags Staff
// @Description Get staff information
// @ModuleID staffInfo
// @Accept json
// @Produce json
// @Success 201 {string} string "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /staff/info [get]
func (h *Handler) staffInfo(c *gin.Context) {
	uuid := c.GetString("uuid")
	if uuid == "" {
		newResponse(c, http.StatusBadRequest, "UUID isn't given")

		return
	}

	res, err := h.services.Staff.StaffInfo(uuid)
	if err != nil {
		if errors.Is(err, model.ErrStaffNotFound) {
			newResponse(c, http.StatusBadRequest, err.Error())

			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Update Staff Info
// @Tags Staff
// @Description Update staff information
// @ModuleID staffUpdateInfo
// @Accept json
// @Produce json
// @Param uuid path string true "Staff uuid"
// @Param input body staffUpdateInfoInput true "Update Staff Information"
// @Success 201 {string} string "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /staff/info/{uuid} [post]
func (h *Handler) staffUpdateInfo(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		newResponse(c, http.StatusBadRequest, "UUID is empty")

		return
	}

	var input staffUpdateInfoInput
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if err := h.services.Staff.StaffUpdateInfo(uuid, model.UpdateStaffInfoInput{
		Name:  input.Name,
		Phone: input.Phone,
	}); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	message := fmt.Sprintf("Staff: %s successfully updated!", input.Name)

	c.JSON(http.StatusOK, map[string]string{
		"message": message,
	})
}
