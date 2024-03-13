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

// Auth client from firebase
var Auth *auth.Client

// Initialize Firebase and its related clients
func Init() {
	firebaseApp, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing Firebase: %v\n", err)
	}

	firebaseAuth, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("error initializing Firebase auth: %v\n", err)
	}

	log.Println("firebase initialized")

	Auth = firebaseAuth
}

func Protect(handlerFunc echo.HandlerFunc, allowedRoles ...string) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header.Get("Authorization")

		if authorization == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing authorization header")
		}

		authorization = strings.Replace(authorization, "Bearer ", "", 1)
		token, err := Auth.VerifyIDToken(context.Background(), authorization)

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
