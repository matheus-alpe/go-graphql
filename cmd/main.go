package main

import (
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/matheus-alpe/go-graphql/internal/database"
	"github.com/matheus-alpe/go-graphql/internal/graphql"
)

func main() {
	db := database.InitDB()

	schema, err := graphql.CreateBlogSchema(db)
	if err != nil {
		log.Fatal("failed to create schema, error:", err)
	}

	http.Handle("/blogs", handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	}))
	http.ListenAndServe(":8080", nil)
}
