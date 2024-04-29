package graphql

import (
	"database/sql"

	"github.com/graphql-go/graphql"
)

func QueryType(blogType *graphql.Object, db *sql.DB) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"blogs": &graphql.Field{
				Type: graphql.NewList(blogType),
				Args: graphql.FieldConfigArgument{
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: GetBlogs(db),
			},
		},
	})
}
