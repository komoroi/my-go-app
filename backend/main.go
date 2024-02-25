package main

import (
	"fmt"
	"log"
	"my-go-app/db"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// ログファイルを作成
	logFile := createLogFile()
	defer logFile.Close()

	// ログの出力先をログファイルに設定
	log.SetOutput(logFile)

	db.NewDB()

	// e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.JSON(200, map[string]string{
	// 		"message": "こんにちは、世界！",
	// 	})
	// })
	// e.Start(":8080")
	log.Println("listening on :8080qqqqq")
	// db.NewDB()
	// echoフレームワークのデフォルトの設定を使用してルータを作成
	// e := echo.New()

	// Create an ent.Client with PostgreSQL database.
	// entOptions := []ent.Option{}
	// entOptions = append(entOptions, ent.Debug())
	// client, err := ent.Open(dialect.Postgres, "host=my-postgres-db user=myuser dbname=mydb sslmode=disable password=mypassword", entOptions...)
	// if err != nil {
	// 	log.Fatalf("failed opening connection to postgres: %v", err)
	// }
	// defer client.Close()

	// // Run the automatic migration tool to create all schema resources.
	// if err := client.Schema.Create(context.Background(), sql.WithGlobalUniqueID(true)); err != nil {
	// 	log.Fatalf("failed creating schema resources: %v", err)
	// }

	// ctx := context.Background()
	// u, err := client.User.
	// 	Create().
	// 	SetName("John").
	// 	SetEmail("hogehoge@gmail.com").
	// 	Save(ctx)
	// if err != nil {
	// 	log.Fatalf("failed creating user: %v", err)
	// }
	// log.Println("user was created: ", u)

	// port := os.Getenv("PORT")

	// srv := handler.NewDefaultServer(resolver.NewSchema(client))

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))
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
