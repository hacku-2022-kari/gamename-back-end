package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// --------------------

	// Rooting

	// --------------------
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", home)

	// Start HTTP server.
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

// e.GET("/", home)
func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Golang + echo!")
}
