package main

import (
	"context"
	"html/template"
	"log"

	"github.com/e9cryptteam/coinman/db"
	"github.com/e9cryptteam/coinman/pkg/handlers"
	"github.com/e9cryptteam/coinman/pkg/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()

	d, err := pgxpool.New(ctx, "user=e9admin dbname=coinman sslmode=disable host=127.0.0.1 port=5432 password=hjkl1234")
	if err != nil {
		panic(err)
	}

	templates, err := template.ParseFiles(
		"public/views/header.html",
		"public/views/coins/index.html",
		"public/views/coins/coins.html",
	)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	q := db.New(d)
	coinHandlers := handlers.NewCoinHandler(*q, ctx)

	e.Renderer = utils.NewTemplateRenderer(templates)
	e.Static("/css", "public/css")

	e.GET("/", coinHandlers.Index)

	e.Logger.Fatal(e.Start(":8080"))
}
