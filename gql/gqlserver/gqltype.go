package gqlserver

import "github.com/graphql-go/graphql"

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "User",
		Description: "Detail of the user",
		Fields: graphql.Fields{
			"user_id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"birth_date": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"phone_number": &graphql.Field{
				Type: graphql.String,
			},
			"membership": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.String,
			},
			"updated_at": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var CouponType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Coupon",
		Description: "Detail of the coupon",
		Fields: graphql.Fields{
			"coupon_id": &graphql.Field{
				Type: graphql.Int,
			},
			"coupon_name": &graphql.Field{
				Type: graphql.String,
			},
			"banner_url": &graphql.Field{
				Type: graphql.String,
			},
			"category": &graphql.Field{
				Type: graphql.String,
			},
			"type": &graphql.Field{
				Type: graphql.String,
			},
			"coupon_percentage": &graphql.Field{
				Type: graphql.Float,
			},
			"coupon_max_nominal": &graphql.Field{
				Type: graphql.Float,
			},
			"min_transaction": &graphql.Field{
				Type: graphql.Float,
			},
			"start_date": &graphql.Field{
				Type: graphql.String,
			},
			"expire_date": &graphql.Field{
				Type: graphql.String,
			},
			"terms_and_condition": &graphql.Field{
				Type: graphql.String,
			},
			"guide": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.String,
			},
			"updated_at": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var UserCouponType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "UserCopon",
		Description: "Detail of the user coupon",
		Fields: graphql.Fields{
			"user_coupon_id": &graphql.Field{
				Type: graphql.Int,
			},
			"user_id": &graphql.Field{
				Type: graphql.Int,
			},
			"coupon_id": &graphql.Field{
				Type: graphql.Int,
			},
			"is_used": &graphql.Field{
				Type: graphql.Boolean,
			},
			"created_at": &graphql.Field{
				Type: graphql.String,
			},
			"updated_at": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var ResultType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Result",
		Description: "Result of operation",
		Fields: graphql.Fields{
			"success": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)
