package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/api/iterator"
)

func main() {
	// --------------------

	// Rooting

	// --------------------
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", home)
	e.POST("/add-user", addUser)

	// Start HTTP server.
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func defineFirebase() (context.Context, *firestore.Client) {
	// --------------------

	// Setting Environment Variables

	// --------------------
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Could not load environment variables: %v", err)
	}
	project_id := os.Getenv("PROJECT_ID")

	// --------------------

	// Firebase

	// --------------------
	// Use the application default credentials
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: project_id}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln("error", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln("error", err)
	}

	return ctx, client
}

// e.GET("/", home)
func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Golang + echo!")
}

// e.POST("/add-user", user)
// curl -d "first=Post" -d "last=From Go" -d "born=1999" http://localhost:8080/add-user
func addUser(c echo.Context) error {
	ctx, client := defineFirebase()

	first := c.FormValue("first")
	last := c.FormValue("last")
	born := c.FormValue("born")
	born_int, _ := strconv.Atoi(born)

	_, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
		"first": first,
		"last":  last,
		"born":  born_int,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	return c.String(http.StatusOK, "Success!")
}
