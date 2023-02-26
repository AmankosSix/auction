package v1

import (
	"auction/internal/model"
	"auction/internal/service"
	"errors"
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
	}
}

// Sign Up
type userSignUpInput struct {
	Name     string `json:"name" binding:"required,min=2,max=64"`
	Email    string `json:"email" binding:"required,email,max=64"`
	Phone    string `json:"phone" binding:"required,max=12"`
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

// @Summary User SignUp
// @Tags Auth
// @Description create user account
// @ModuleID userSignUp
// @Accept json
// @Produce json
// @Param input body userSignUpInput true "sign up info"
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

	c.Status(http.StatusCreated)
}

// @Summary User SignIn
// @Tags Auth
// @Description log in user account
// @ModuleID userSignIn
// @Accept json
// @Produce json
// @Param input body userSignInInput true "sign in info"
// @Success 200 {object} userSignInResponse
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
// @Tags Auth
// @Description user info
// @ModuleID userInfo
// @Accept json
// @Produce json
// @Param input body userSignInInput true "sign in info"
// @Success 200 {object} model.UserInfo
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
