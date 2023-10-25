package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(string)
					//Vai printar a id passada na query
					//Ex.: query={user(id:"1"){id, name}}
					fmt.Println(idQuery)
					if isOK {
						//Buscar do banco de dados e
						//retornar um map do tipo usuário
						return idQuery, nil
					}
					return nil, nil
				},
			},
		},
	},
)
