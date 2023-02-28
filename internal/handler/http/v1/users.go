package v1

import (
	"auction/internal/model"
	"auction/internal/service"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	auth := api.Group("/auth")
	{
		auth.POST("/sign-up", h.userSignUp)
		auth.POST("/sign-in", h.userSignIn)
		auth.POST("/refresh", h.userSignIn)
	}
	user := api.Group("/user", h.userIdentity)
	{
		user.GET("/info", h.userInfo)
		user.POST("/info/:uuid", h.userUpdateInfo)
	}
}

// Sign Up
type userSignUpInput struct {
	Name     string `json:"name" binding:"required,min=2,max=64"`
	Email    string `json:"email" binding:"required,email,max=64"`
	Phone    string `json:"phone" binding:"required,max=10"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

type userSignInInput struct {
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"'`
}

// Sign In
type userSignInResponse struct {
	AccessToken string `json:"accessToken"`
}

// User Info Update
type userUpdateInfoInput struct {
	Name  string `json:"name" binding:"required,min=2,max=64"`
	Phone string `json:"phone" binding:"required,max=10"`
}

// @Summary Sign up
// @Tags Account
// @Description Create account with role user
// @ModuleID userSignUp
// @Accept json
// @Produce json
// @Param input body userSignUpInput true "Sign Up"
// @Success 201 {string} string "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) userSignUp(c *gin.Context) {
	var input userSignUpInput
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if err := h.services.Users.SignUp(c.Request.Context(), service.UserSignUpInput{
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
		Password: input.Password,
	}); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	message := fmt.Sprintf("User: %s successfully created!", input.Name)

	c.JSON(http.StatusCreated, map[string]string{
		"message": message,
	})
}

// @Summary Sign in
// @Tags Account
// @Description Sign in to user account
// @ModuleID userSignIn
// @Accept json
// @Produce json
// @Param input body userSignInInput true "Sign In"
// @Success 201 {string} string "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) userSignIn(c *gin.Context) {
	var input userSignInInput
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")

		return
	}

	res, err := h.services.Users.SignIn(c.Request.Context(), service.UserSignInInput{
		Email:    input.Email,
		Password: input.Password,
	})

	if err != nil {
		if errors.Is(err, model.ErrUserNotFound) {
			newResponse(c, http.StatusBadRequest, err.Error())

			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.SetCookie("refresh_token", res.RefreshToken, 60*60, "/auth", "localhost", false, true)

	c.JSON(http.StatusOK, userSignInResponse{
		AccessToken: res.AccessToken,
	})
}

// @Summary User Info
// @Tags Account
// @Description Get user information
// @ModuleID userInfo
// @Accept json
// @Produce json
// @Success 201 {string} string "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/info [get]
func (h *Handler) userInfo(c *gin.Context) {
	uuid := c.GetString("uuid")
	if uuid == "" {
		newResponse(c, http.StatusBadRequest, "UUID isn't given")

		return
	}

	res, err := h.services.Users.UserInfo(uuid)
	if err != nil {
		if errors.Is(err, model.ErrUserNotFound) {
			newResponse(c, http.StatusBadRequest, err.Error())

			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Update User Info
// @Tags Account
// @Description Update user information
// @ModuleID userUpdateInfo
// @Accept json
// @Produce json
// @Param uuid path string true "User uuid"
// @Param input body userUpdateInfoInput true "Update User Information"
// @Success 201 {string} string "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/info/{uuid} [post]
func (h *Handler) userUpdateInfo(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		newResponse(c, http.StatusBadRequest, "UUID is empty")

		return
	}

	var input userUpdateInfoInput
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if err := h.services.Users.UserUpdateInfo(uuid, model.UpdateUserInfoInput{
		Name:  input.Name,
		Phone: input.Phone,
	}); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	message := fmt.Sprintf("User: %s successfully updated!", input.Name)

	c.JSON(http.StatusOK, map[string]string{
		"message": message,
	})
}
