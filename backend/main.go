package main

import (
	"fmt"
	"log"
	"my-go-app/ent"
	"my-go-app/resolver"
	"net/http"
	"os"
	"time"

	// "github.com/onikan27/graphql-test-api/graph"
	// "my-go-app/resolver/model"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	// ログファイルを作成
	logFile := createLogFile()
	defer logFile.Close()

	// ログの出力先をログファイルに設定
	log.SetOutput(logFile)

	// Create an ent.Client with PostgreSQL database.
	entOptions := []ent.Option{}
	entOptions = append(entOptions, ent.Debug())

	// .env ファイルを読み込む
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	// .env ファイルから必要な情報を取得
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	// PostgreSQL 接続文字列を構築
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword)
	// ent フレームワークを使用して PostgreSQL に接続
	client, err := ent.Open(dialect.Postgres, connStr)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer client.Close()

	log.Println(resolver.NewSchema(client))

	srv := handler.NewDefaultServer(resolver.NewSchema(client))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"},
		AllowCredentials: true,
	})

	log.Println(c)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", c.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func createLogFile() *os.File {
	// Set the location to Japan Standard Time (JST)
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalf("タイムゾーンの設定に失敗しました: %v", err)
	}

	// ログディレクトリを作成
	logDir := "logs"
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatalf("ログディレクトリの作成に失敗しました: %v", err)
	}

	// ログファイルの名前を年月日で作成（日本時間）
	logFileName := fmt.Sprintf("%s/%s.log", logDir, time.Now().In(jst).Format("2006-01-02"))

	// ログファイルを書き込み用に作成またはオープン
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("ログファイルのオープンに失敗しました: %v", err)
	}

	return logFile
}
