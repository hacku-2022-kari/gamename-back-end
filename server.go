package main

import (
	"context"
	"encoding/json"
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

type User struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Born  int    `json:"born"`
}

func main() {
	// --------------------

	// Rooting

	// --------------------
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", home)
	e.GET("/user-list", userList)
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

// e.GET("/user-list", userList)
func userList(c echo.Context) error {
	ctx, client := defineFirebase()

	iter := client.Collection("users").Documents(ctx)
	var output []User

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		bytes, err := json.Marshal(doc.Data())
		if err != nil {
			fmt.Println("JSON marshal error: ", err)
			return err
		}

		var user User
		err = json.Unmarshal(bytes, &user)
		if err != nil {
			return err
		}
		output = append(output, user)
	}

	return c.JSON(http.StatusOK, output)
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
