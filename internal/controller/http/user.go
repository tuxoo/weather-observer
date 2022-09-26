package http

import "github.com/gin-gonic/gin"

func (h *Handler) initUserRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/sign-up", h.signUp)
		users.POST("/sign-in", h.signIn)
	}
}

func (h *Handler) signUp(c *gin.Context) {

}

func (h *Handler) signIn(c *gin.Context) {

}
