package main

import (
	"context"
	"html/template"
	"log"

	"github.com/e9cryptteam/coinman/db"
	"github.com/e9cryptteam/coinman/pkg/api"
	"github.com/e9cryptteam/coinman/pkg/utils"
	"github.com/e9cryptteam/coinman/pkg/web"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const ViewDir = "public/views"

func main() {
	ctx := context.Background()

	d, err := pgxpool.New(ctx, "user=e9admin dbname=coinman sslmode=disable host=127.0.0.1 port=5432 password=hjkl1234")
	if err != nil {
		panic(err)
	}

	/*templates, err := template.ParseFiles(
		"public/views/header.html",
		"public/views/nav.html",
		"public/views/coins/index.html",
		"public/views/coins/coins.html",
		"public/views/markets/market-table.html",
		"public/views/markets/markets.html",
	)
	if err != nil {
		log.Fatal(err)
	}*/
	templates, err := template.ParseGlob("public/views/*.html")
  templates.ParseGlob("public/views/coins/*.html")
  templates.ParseGlob("public/views/markets/*.html")
	if err != nil {
		log.Fatal(err)
	}

	for _, t := range templates.Templates() {
		log.Println(t.Name())
	}

	conn := db.New(d)

	e := echo.New()
	e.Renderer = utils.NewTemplateRenderer(templates)
	e.Debug = true

	e.Use(middleware.Logger())

	e.Static("/css", "public/css")
	e.Static("/js", "public/js")

	api.InitRoutes(e, *conn, ctx)
	web.InitRoutes(e, *conn, ctx)

	e.Logger.Fatal(e.Start(":8080"))
}
