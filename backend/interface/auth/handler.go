package auth

import (
	"backend/application/auth"
	"backend/sorry"

	"net/http"

	"github.com/gin-gonic/gin"
)

// login godoc
// @Summary Login
// @Description Login as an admin
// @Tags auth
// @Accept  json
// @Produce  json
// @Param input body auth.LoginInput true "login input"
// @Success 200 {object} auth.LoginOutput
// @Router /auth/login [post]
func login(c *gin.Context) {
	var (
		credentirals = new(auth.LoginInput)
		token        *auth.LoginOutput
		err          error
	)

	if err = c.ShouldBindJSON(credentirals); err != nil {
		sorry.Handling(c, err)
		return
	}

	if token, err = auth.Login(credentirals); err != nil {
		sorry.Handling(c, err)
		return
	}

	c.JSON(http.StatusOK, token)
}
