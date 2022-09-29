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

// @Summary		User SignUp
// @Tags        authentication
// @Description registers a new user
// @ID          userSignUp
// @Accept      json
// @Produce     json
// @Param       input body dto.SignUpDTO  true  "account information"
// @Success     201
// @Failure     400  	  		{object}  errorResponse
// @Failure     500      		{object}  errorResponse
// @Failure     default  		{object}  errorResponse
// @Router      /users/sign-up 	[post]
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

// @Summary 	User SignIn
// @Tags 		authentication
// @Description authenticates the user
// @ID 			userSignIn
// @Accept  	json
// @Produce  	json
// @Param input body dto.SignInDTO true "credentials"
// @Success 	200 {object} 		dto.LoginResponse
// @Failure 	400,404 {object} 	errorResponse
// @Failure 	500 {object} 		errorResponse
// @Failure 	default {object} 	errorResponse
// @Router 		/users/sign-in 		[post]
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

// @Summary 	User Profile
// @Security 	Bearer
// @Tags 		user
// @Description gets current profile user
// @ID 			currentUser
// @Accept  	json
// @Produce  	json
// @Success 	200 {object} 		dto.User
// @Failure 	500 {object} 		errorResponse
// @Failure 	default {object} 	errorResponse
// @Router 		/users/profile 		[get]
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
