package db

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"my-go-app/ent"
	"my-go-app/ent/migrate"
	"my-go-app/ent/user"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

func NewDB() {
	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve database connection information from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	// Create the database connection string
	dbConnectionString := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " dbname=" + dbName + " password=" + dbPassword + " sslmode=" + dbSSLMode

	// Open the database connection
	client, err := ent.Open("postgres", dbConnectionString)

	// client, err := ent.Open(dialect.Postgres, "host=db user=postgres dbname=db sslmode=disable password=mypassword")
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	// ctx := context.Background()
	// u, err := client.User.
	// 	Create().
	// 	SetName("John").
	// 	SetEmail("hoge@gmail.com").
	// 	SetPassword("password").
	// 	// SetDone(false).
	// 	Save(ctx)
	// if err != nil {
	// 	log.Fatalf("failed creating user: %v", err)
	// }
	// log.Println("user was created: ", u)

	// Configure the server and start listening on :8081.
	// srv := handler.NewDefaultServer(resolver.NewSchema(client))
	// http.Handle("/",
	// 	playground.Handler("Todo", "/query"),
	// )
	// http.Handle("/query", srv)
	// fmt.Println("aaa")
	e := echo.New()
	// fmt.Println(client.User.Query().Where(user.ID(4294967296)).Only(context.Background()))
	e.GET("/", func(c echo.Context) error {
		fmt.Println("u")
		// Query User table for user with ID 1
		u, err := client.User.Query().Where(user.ID(4294967296)).Only(context.Background())
		fmt.Println(u)
		fmt.Println(u)
		if err != nil {
			log.Printf("failed to query user: %v", err)
			return c.String(http.StatusInternalServerError, "Internal Server Error")
		}
		fmt.Println(u)
		// Display user information
		return c.String(http.StatusOK, fmt.Sprintf("User ID: %d, Name: %s, Email: %s", u.ID, u.Name, u.Email))
	})
	e.Start(":8080")
	log.Println("listening on :8080")
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal("http server terminated", err)
	// }

	// Your code here...
}
