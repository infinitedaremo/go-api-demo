package main

import (
	"context"
	"database/sql"

	dbembed "github.com/infinitedaremo/go-api-demo/db"
	"github.com/infinitedaremo/go-api-demo/internal/app"
	db2 "github.com/infinitedaremo/go-api-demo/internal/db"
	"github.com/infinitedaremo/go-api-demo/internal/server"
	"go.uber.org/zap"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx := context.Background()

	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	// Open DB
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal("failed creating db", zap.Error(err))
	}

	// Create tables
	if _, err = db.ExecContext(ctx, dbembed.TableSchema); err != nil {
		log.Fatal("failed creating tables", zap.Error(err))
	}

	queries := db2.New(db)

	// Bootstrap the app
	if err = app.Bootstrap(ctx, queries); err != nil {
		log.Fatal("failed bootstrapping app", zap.Error(err))
	}

	s, err := server.NewServer(log, app.NewPersonService(queries))
	if err != nil {
		log.Fatal("failed creating server", zap.Error(err))
	}

	log.Info("launching server")
	if err = s.Serve(":8080"); err != nil {
		log.Fatal("failed running server", zap.Error(err))
	}
}
