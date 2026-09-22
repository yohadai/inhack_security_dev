package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func jwtSecret() []byte {
	s := os.Getenv("JWT_SECRET")
	if s == "" {
		s = "hs256-devkey-GuOd"
	}
	return []byte(s)
}

func authServiceURL() string {
	u := os.Getenv("AUTH_SERVICE_URL")
	if u == "" {
		u = "http://localhost:8081"
	}
	return u
}

// Authorization 헤더의 Bearer 토큰을 검증하고 사용자 이름을 돌려준다
func verifyToken(c echo.Context) (string, error) {
	header := c.Request().Header.Get("Authorization")
	raw := strings.TrimPrefix(header, "Bearer ")
	token, err := jwt.Parse(raw, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret(), nil
	})
	if err != nil || !token.Valid {
		return "", echo.ErrUnauthorized
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", echo.ErrUnauthorized
	}
	sub, _ := claims["sub"].(string)
	return sub, nil
}

func me(c echo.Context) error {
	username, err := verifyToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
	}
	return c.JSON(http.StatusOK, echo.Map{"username": username})
}

// api-service 가 auth-service 를 HTTP 로 호출해서 상태를 함께 확인한다
func health(c echo.Context) error {
	authStatus := "unknown"
	client := http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get(authServiceURL() + "/health")
	if err == nil {
		defer resp.Body.Close()
		var body map[string]interface{}
		if json.NewDecoder(resp.Body).Decode(&body) == nil {
			authStatus = "reachable"
		}
	} else {
		authStatus = "unreachable"
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status":       "ok",
		"service":      "api",
		"auth_service": authStatus,
	})
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", health)
	e.GET("/me", me)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
