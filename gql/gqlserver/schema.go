package gqlserver

import (
	"github.com/graphql-go/graphql"
)

type SchemaWrapper struct {
	couponResolver *Resolver
	Schema         graphql.Schema
}

func NewSchemaWrapper() *SchemaWrapper {
	return &SchemaWrapper{}
}

func (s *SchemaWrapper) WithCouponResolver(pr *Resolver) *SchemaWrapper {
	s.couponResolver = pr

	return s
}

func (s *SchemaWrapper) Init() error {
	// add gql schema as needed
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:        "UserGetter",
			Description: "All query related to getting user data",
			Fields: graphql.Fields{
				"UserDetail": &graphql.Field{
					Type:        UserType,
					Description: "Get user by ID",
					Args: graphql.FieldConfigArgument{
						"user_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: s.couponResolver.GetUser(),
				},
				"User": &graphql.Field{
					Type:        graphql.NewList(UserType),
					Description: "Get users",
					Resolve:     s.couponResolver.GetUsers(),
				},
				"CouponDetail": &graphql.Field{
					Type:        CouponType,
					Description: "Get coupon by ID",
					Args: graphql.FieldConfigArgument{
						"coupon_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: s.couponResolver.GetCoupon(),
				},
				"Coupon": &graphql.Field{
					Type:        graphql.NewList(CouponType),
					Description: "Get coupons",
					Resolve:     s.couponResolver.GetCoupons(),
				},
			},
		}),

		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:        "UserCouponSetter",
			Description: "All query related to modify coupon data",
			Fields: graphql.Fields{
				"CreateCoupon": &graphql.Field{
					Type:        graphql.Boolean,
					Description: "Create coupon",
					Args: graphql.FieldConfigArgument{
						"coupon_name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"banner_url": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"category": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"type": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"coupon_percentage": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
						"coupon_max_nominal": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
						"min_transaction": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
						"start_date": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"expire_date": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"terms_and_condition": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"guide": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"created_at": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"updated_at": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: s.couponResolver.CreateCoupons(),
				},
				"UpdateCoupon": &graphql.Field{
					Type:        graphql.Boolean,
					Description: "Update coupon by ID",
					Args: graphql.FieldConfigArgument{
						"coupon_id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"coupon_name": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
						"banner_url": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"category": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"type": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"coupon_percentage": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
						"coupon_max_nominal": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
						"min_transaction": &graphql.ArgumentConfig{
							Type: graphql.Float,
						},
						"start_date": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"expire_date": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"terms_and_condition": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"guide": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"created_at": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
						"updated_at": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: s.couponResolver.UpdateCoupons(),
				},
			},
		}),

		// // uncomment this and add objects for mutation
		// Mutation: graphql.NewObject(graphql.ObjectConfig{
		// 	Name:        "ProductSetter",
		// 	Description: "All query related to modify product data",
		// 	Fields: graphql.Fields{
		// 		"CreateProduct": &graphql.Field{
		// 			Type:        graphql.Boolean,
		// 			Description: "Create product",
		// 			Args: graphql.FieldConfigArgument{
		// 				"product_name": &graphql.ArgumentConfig{
		// 					Type: graphql.NewNonNull(graphql.String),
		// 				},
		// 				"product_shop_name": &graphql.ArgumentConfig{
		// 					Type: graphql.String,
		// 				},
		// 				"product_price": &graphql.ArgumentConfig{
		// 					Type: graphql.Float,
		// 				},
		// 				"product_image": &graphql.ArgumentConfig{
		// 					Type: graphql.String,
		// 				},
		// 			},
		// 			Resolve: s.productResolver.CreateProducts(),
		// 		},
		// 		"UpdateProduct": &graphql.Field{
		// 			Type:        graphql.Boolean,
		// 			Description: "Update product by ID",
		// 			Args: graphql.FieldConfigArgument{
		// 				"product_id": &graphql.ArgumentConfig{
		// 					Type: graphql.NewNonNull(graphql.Int),
		// 				},
		// 				"product_name": &graphql.ArgumentConfig{
		// 					Type: graphql.NewNonNull(graphql.String),
		// 				},
		// 				"product_shop_name": &graphql.ArgumentConfig{
		// 					Type: graphql.String,
		// 				},
		// 				"product_price": &graphql.ArgumentConfig{
		// 					Type: graphql.Float,
		// 				},
		// 				"product_image": &graphql.ArgumentConfig{
		// 					Type: graphql.String,
		// 				},
		// 			},
		// 			Resolve: s.productResolver.UpdateProducts(),
		// 		},
		// 	},
		// }),
	})

	if err != nil {
		return err
	}

	s.Schema = schema

	return nil
}
