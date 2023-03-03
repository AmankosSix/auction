package v1

import (
	"auction/internal/model"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initDictionariesRoutes(api *gin.RouterGroup) {
	dictionaries := api.Group("/dictionaries", h.userIdentity)
	{
		dictionaries.GET("/roles", h.rolesList)
	}
}

// @Summary Dictionaries
// @Tags Dictionaries
// @Description Get all roles
// @ModuleID rolesList
// @Accept json
// @Produce json
// @Success 201 {string} string "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /owner/staff/list [get]
func (h *Handler) rolesList(c *gin.Context) {
	res, err := h.services.Dictionaries.RolesList()
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
