package resolvers

import "github.com/graphql-go/graphql"

func GetAllCoupon() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		return "All Coupon",nil
	}
}