package graphql

import (
	"database/sql"
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/matheus-alpe/go-graphql/internal/model"
)

func createBlogType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Blog",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"content": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

func CreateBlogSchema(db *sql.DB) (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: QueryType(createBlogType(), db),
	})
}

func GetBlogs(db *sql.DB) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		limit, _ := p.Args["limit"].(int)
		if limit <= 0 || limit > 20 {
			limit = 10
		}

		offset, _ := p.Args["offset"].(int)
		if offset < 0 {
			offset = 0
		}

		query := fmt.Sprintf("SELECT id, title, content FROM blogs limit %v offset %v", limit, offset)

		var blogs []model.Blog
		rows, err := db.Query(query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var b model.Blog
			if err := rows.Scan(&b.ID, &b.Title, &b.Content); err != nil {
				return nil, err
			}

			blogs = append(blogs, b)
		}

		return blogs, nil
	}
}
