package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tuxoo/weather-observer/internal/model/dto"
	"net/http"
)

func (h *Handler) initUserRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/sign-up", h.signUp)
		users.POST("/sign-in", h.signIn)

		authenticated := users.Group("/", h.userIdentity)
		{
			authenticated.GET("/profile", h.getUserProfile)
		}
	}
}

func (h *Handler) signUp(ctx *gin.Context) {
	var signUpDto dto.SignUpDTO
	if err := ctx.ShouldBindJSON(&signUpDto); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.userService.SignUp(ctx.Request.Context(), signUpDto); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusCreated)
}

func (h *Handler) signIn(ctx *gin.Context) {
	var signInDTO dto.SignInDTO
	if err := ctx.ShouldBindJSON(&signInDTO); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	loginResponse, err := h.userService.SignIn(ctx.Request.Context(), signInDTO)
	if err != nil {
		newErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, loginResponse)
}

func (h *Handler) getUserProfile(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.userService.GetById(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.User{
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		RegisteredAt: user.RegisteredAt,
	})
}
