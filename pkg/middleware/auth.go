package middleware

import (
	resultdto "landtick/dto/result"
	jwToken "landtick/pkg/jwt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func Auth(Next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return c.JSON(http.StatusUnauthorized, resultdto.ErrorResult{Status: "Failed", Message: "Unauthorized"})
		}

		token = strings.Split(token, " ")[1]
		claims, err := jwToken.DecodeToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, resultdto.ErrorResult{Status: "Failed", Message: "Unauthorized"})
		}

		c.Set("userLogin", claims)
		return Next(c)
	}
}
