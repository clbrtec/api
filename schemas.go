package main

import (
	"github.com/graphql-go/graphql"
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)
