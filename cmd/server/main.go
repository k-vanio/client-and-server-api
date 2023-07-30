package main

import (
	"net/http"
	"time"

	"github.com/x-vanio/client-and-server-api/internal/db"
	"github.com/x-vanio/client-and-server-api/internal/model"
	"github.com/x-vanio/client-and-server-api/pkg/request"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	conn, err := gorm.Open(sqlite.Open("currency.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	conn.AutoMigrate(&model.Quote{})

	client := request.NewClient(http.DefaultClient, time.Millisecond*200)
	repository := &db.SqliteRepository{
		DB: conn,
	}

	// mux
	mux := http.NewServeMux()

	// handler
	handler := NewHandler(client, repository)
	mux.HandleFunc("/cotacao", handler.GetDollarQuote)

	// server
	server := NewServer(mux, 8080)
	server.Start()
}
