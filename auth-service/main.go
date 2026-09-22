package main

import (
	"net/http"
	"os"
	"sync"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// 메모리에 저장하는 사용자 한 명의 정보
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 아주 단순한 인메모리 사용자 저장소 (DB 대신 사용)
var (
	users   = map[string]User{}
	usersMu sync.Mutex
)

func jwtSecret() []byte {
	s := os.Getenv("JWT_SECRET")
	if s == "" {
		s = "hs256-devkey-GuOd"
	}
	return []byte(s)
}

func signup(c echo.Context) error {
	var u User
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid body"})
	}
	if u.Username == "" || u.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "username and password required"})
	}

	usersMu.Lock()
	defer usersMu.Unlock()
	if _, exists := users[u.Username]; exists {
		return c.JSON(http.StatusConflict, echo.Map{"error": "user already exists"})
	}
	users[u.Username] = u
	return c.JSON(http.StatusCreated, echo.Map{"message": "signup ok"})
}

func login(c echo.Context) error {
	var u User
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid body"})
	}

	usersMu.Lock()
	saved, ok := users[u.Username]
	usersMu.Unlock()
	if !ok || saved.Password != u.Password {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid credentials"})
	}

	// 로그인 성공 시 JWT 토큰을 발급한다
	claims := jwt.MapClaims{
		"sub": u.Username,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(jwtSecret())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "token error"})
	}
	return c.JSON(http.StatusOK, echo.Map{"token": signed})
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"status": "ok", "service": "auth"})
	})
	e.POST("/signup", signup)
	e.POST("/login", login)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
