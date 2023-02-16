package v1

import (
	"auction/internal/model"
	"auction/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/sign-up", h.userSignUp)
		users.POST("/sign-in", h.userSignIn)
	}
}

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

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// @Summary User SignUp
// @Tags users-auth
// @Description create user account
// @ModuleID userSignUp
// @Accept json
// @Produce json
// @Param input body userSignUpInput true "sign up info"
// @Success 201 {string} string "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/sign-up [post]
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

// @Summary User SignIp
// @Tags users-auth
// @Description log in user account
// @ModuleID userSignIn
// @Accept json
// @Produce json
// @Param input body userSignInInput true "sign in info"
// @Success 200 {object} tokenResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/sign-in [post]
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

	c.JSON(http.StatusOK, tokenResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	})
}
