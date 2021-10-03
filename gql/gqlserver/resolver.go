package gqlserver

import (
	"github.com/graphql-go/graphql"
	"github.com/peterchu999/tgtc/backend/dictionary"
	"github.com/peterchu999/tgtc/backend/service"
)

type Resolver struct {
}

func NewResolver() *Resolver {
	return &Resolver{}
}

func (r *Resolver) GetUser() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["user_id"].(int)

		// update to use Usecase from previous session
		return service.GetUser(id)
	}
}

func (r *Resolver) GetUsers() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		return service.GetUsers()
	}
}

func (r *Resolver) GetCoupon() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["coupon_id"].(int)

		// update to use Usecase from previous session
		return service.GetCoupon(id)
	}
}

func (r *Resolver) GetCoupons() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		return service.GetCoupons()
	}
}

func (r *Resolver) CreateCoupons() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		couponName, _ := p.Args["coupon_name"].(string)
		bannerUrl, _ := p.Args["banner_url"].(string)
		category, _ := p.Args["category"].(string)
		type, _ := p.Args["type"].(string)
		couponPercentage, _ := p.Args["coupon_percentage"].(string)
		couponMaxNominal, _ := p.Args["coupon_max_nominal"].(float64)
		minTransaction, _ := p.Args["min_transaction"].(float64)
		startDate, _ := p.Args["start_date"].(string)
		expireDate, _ := p.Args["expire_date"].(string)
		termsCondition, _ := p.Args["terms_and_condition"].(string)
		guide, _ := p.Args["guide"].(string)
		createdAt, _ := p.Args["created_at"].(string)
		updatedAt, _ := p.Args["updated_at"].(string)

		req := dictionary.Coupon{
			CouponName: couponName,
			BannerUrl: bannerUrl,
			Category: category,
			Type: type,
			CouponPercentage: couponPercentage,
			CouponMaxNominal: couponMaxNominal,
			MinTransaction: minTransaction,
			StartDate: startDate,
			ExpireDate: expireDate,
			TermsAndCondition: termsCondition,
			Guide: guide,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		_, err := service.CreateCoupon(req)
		if err != nil {
			return false, err
		}

		return true, nil
	}
}

func (r *Resolver) UpdateCoupons() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, _ := p.Args["coupon_id"].(int)
		couponName, _ := p.Args["coupon_name"].(string)
		bannerUrl, _ := p.Args["banner_url"].(string)
		category, _ := p.Args["category"].(string)
		type, _ := p.Args["type"].(string)
		couponPercentage, _ := p.Args["coupon_percentage"].(string)
		couponMaxNominal, _ := p.Args["coupon_max_nominal"].(float64)
		minTransaction, _ := p.Args["min_transaction"].(float64)
		startDate, _ := p.Args["start_date"].(string)
		expireDate, _ := p.Args["expire_date"].(string)
		termsCondition, _ := p.Args["terms_and_condition"].(string)
		guide, _ := p.Args["guide"].(string)
		createdAt, _ := p.Args["created_at"].(string)
		updatedAt, _ := p.Args["updated_at"].(string)

		req := dictionary.Coupon{
			ID:int64(id),
			CouponName: couponName,
			BannerUrl: bannerUrl,
			Category: category,
			Type: type,
			CouponPercentage: couponPercentage,
			CouponMaxNominal: couponMaxNominal,
			MinTransaction: minTransaction,
			StartDate: startDate,
			ExpireDate: expireDate,
			TermsAndCondition: termsCondition,
			Guide: guide,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		_, err := service.UpdateCoupon(req)
		if err != nil {
			return false, err
		}

		return true, nil
	}
}

func (r *Resolver) CreateCoupons() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		couponName, _ := p.Args["coupon_name"].(string)
		bannerUrl, _ := p.Args["banner_url"].(string)
		category, _ := p.Args["category"].(string)
		type, _ := p.Args["type"].(string)
		couponPercentage, _ := p.Args["coupon_percentage"].(string)
		couponMaxNominal, _ := p.Args["coupon_max_nominal"].(float64)
		minTransaction, _ := p.Args["min_transaction"].(float64)
		startDate, _ := p.Args["start_date"].(string)
		expireDate, _ := p.Args["expire_date"].(string)
		termsCondition, _ := p.Args["terms_and_condition"].(string)
		guide, _ := p.Args["guide"].(string)
		createdAt, _ := p.Args["created_at"].(string)
		updatedAt, _ := p.Args["updated_at"].(string)

		req := dictionary.Coupon{
			CouponName: couponName,
			BannerUrl: bannerUrl,
			Category: category,
			Type: type,
			CouponPercentage: couponPercentage,
			CouponMaxNominal: couponMaxNominal,
			MinTransaction: minTransaction,
			StartDate: startDate,
			ExpireDate: expireDate,
			TermsAndCondition: termsCondition,
			Guide: guide,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		_, err := service.CreateCoupon(req)
		if err != nil {
			return false, err
		}

		return true, nil
	}
}

func (r *Resolver) CreateUserCoupons() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		user_id, _ := p.Args["user_id"].(int)
		coupon_id, _ := p.Args["banner_url"].(string)
		is_used, _ := p.Args["is_used"].(string)
		createdAt, _ := p.Args["created_at"].(string)
		updatedAt, _ := p.Args["updated_at"].(string)

		req := dictionary.Coupon{
			UserId: user_id,
			CouponId: coupon_id,
			IsUsed: is_used,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		_, err := service.CreateUserCoupons(req)
		if err != nil {
			return false, err
		}

		return true, nil
	}
}