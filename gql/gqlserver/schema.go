package gqlserver

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/peterchu999/tgtc-project/gql/gqlserver/resolvers"
)


var queryType = graphql.NewObject(graphql.ObjectConfig{
    Name: "Query",
    Fields: graphql.Fields{
        "query1": &graphql.Field{
            Type: graphql.String,
			Args: graphql.FieldConfigArgument{},
            Resolve: resolvers.GetAllCoupon(),
        },
    },
})


func Init() graphql.Schema {
	var Schema, err = graphql.NewSchema(graphql.SchemaConfig{
    Query: queryType,
	})
	fmt.Print(err)
	return Schema
}

