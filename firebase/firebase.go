package firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
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

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing authorization header")
		}

		token, err := Auth.VerifyIDToken(context.Background(), authHeader)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization token")
		}

		c.Set("uid", token.UID)

		return next(c)
	}
}
