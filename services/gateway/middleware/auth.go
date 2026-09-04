package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

const (
	HeaderUserID   = "X-User-ID"
	HeaderUserRole = "X-User-Role"
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func StripClientIdentity(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Request().Header.Del(HeaderUserID)
		c.Request().Header.Del(HeaderUserRole)
		return next(c)
	}
}

func Auth(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			raw := c.Request().Header.Get(echo.HeaderAuthorization)
			if !strings.HasPrefix(raw, "Bearer ") {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing bearer token")
			}

			claims := &Claims{}
			token, err := jwt.ParseWithClaims(
				strings.TrimPrefix(raw, "Bearer "),
				claims,
				func(t *jwt.Token) (any, error) {
					if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, jwt.ErrSignatureInvalid
					}
					return []byte(secret), nil
				},
			)
			if err != nil || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired token")
			}
			c.Request().Header.Set(HeaderUserID, strconv.FormatUint(uint64(claims.UserID), 10))
			c.Request().Header.Set(HeaderUserRole, claims.Role)

			c.Set("user_id", claims.UserID)
			c.Set("role", claims.Role)
			return next(c)
		}
	}
}

func RequireAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Get("role") != "admin" {
			return echo.NewHTTPError(http.StatusForbidden, "admin access required")
		}
		return next(c)
	}
}
