package v1

import (
	"auction/internal/model"
	"auction/internal/service"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initOwnerRoutes(api *gin.RouterGroup) {
	owner := api.Group("/owner/staff", h.ownerIdentity)
	{
		owner.POST("/sign-up", h.staffSignUp)
		owner.GET("/list", h.staffList)
		owner.DELETE("/:uuid", h.removeStaff)
	}
}

// Sign Up
type ownerSignUpInput struct {
	Name     string `json:"name" binding:"required,min=2,max=64"`
	Email    string `json:"email" binding:"required,email,max=64"`
	Phone    string `json:"phone" binding:"required,max=10"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

// @Summary Sign up new staff
// @Tags Owner
// @Description Owner creates staff
// @ModuleID staffSignUp
// @Accept json
// @Produce json
// @Param input body ownerSignUpInput true "Sign Up Staff"
// @Success 201 {string} string "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /owner/staff/sign-up [post]
func (h *Handler) staffSignUp(c *gin.Context) {
	var input ownerSignUpInput
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if err := h.services.Owner.SignUp(c.Request.Context(), service.OwnerSignUpInput{
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
		Password: input.Password,
	}); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	message := fmt.Sprintf("Owner: %s successfully created!", input.Name)

	c.JSON(http.StatusCreated, map[string]string{
		"message": message,
	})
}

// @Summary Staff Info
// @Tags Owner
// @Description Get all staff information
// @ModuleID staffList
// @Accept json
// @Produce json
// @Success 201 {string} string "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /owner/staff/list [get]
func (h *Handler) staffList(c *gin.Context) {
	res, err := h.services.Owner.StaffList()
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

// @Summary Remove Staff
// @Tags Owner
// @Description Remove staff by ID
// @ModuleID removeStaff
// @Accept json
// @Produce json
// @Success 201 {string} string "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /owner/staff/{uuid} [delete]
func (h *Handler) removeStaff(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		newResponse(c, http.StatusBadRequest, "UUID is empty")

		return
	}

	if err := h.services.Owner.RemoveStaff(uuid); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "Staff successfully removed",
	})
}
