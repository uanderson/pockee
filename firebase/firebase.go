package firebase

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
)

type Firebase struct {
	Auth *auth.Client
}

func New() *Firebase {
	firebaseApp, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing Firebase: %v\n", err)
	}

	firebaseAuth, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("error initializing Firebase auth: %v\n", err)
	}

	log.Println("firebase initialized")

	return &Firebase{
		Auth: firebaseAuth,
	}
}

func (f *Firebase) Protect(handlerFunc echo.HandlerFunc, allowedRoles ...string) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header.Get("Authorization")

		if authorization == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing authorization header")
		}

		authorization = strings.Replace(authorization, "Bearer ", "", 1)
		token, err := f.Auth.VerifyIDToken(context.Background(), authorization)

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		c.Set("uid", token.UID)

		if len(allowedRoles) > 0 && !hasRole(allowedRoles, token) {
			return echo.NewHTTPError(http.StatusForbidden, "Access denied")
		}

		return handlerFunc(c)
	}
}

func hasRole(allowedRoles []string, token *auth.Token) bool {
	for _, role := range allowedRoles {
		if _, exists := token.Claims["roles"].(map[string]interface{})[role]; exists {
			return true
		}
	}
	return false
}
