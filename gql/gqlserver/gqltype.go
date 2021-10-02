package gqlserver

import "github.com/graphql-go/graphql"

var CouponType = graphql.NewObject(
	graphql.ObjectConfig {
		Name: "Coupon",
		Description: "Coupon Description",
		Fields: graphql.Fields {
			"field1" : &graphql.Field{
				Type: graphql.String,
			},
		},
	},
) 

var UserType = graphql.NewObject(
	graphql.ObjectConfig {
		Name: "User",
		Description: "User Description",
		Fields: graphql.Fields {
			"field1" : &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)