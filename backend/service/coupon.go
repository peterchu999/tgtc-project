package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/peterchu999/tgtc/backend/database"
	"github.com/peterchu999/tgtc/backend/dictionary"
)

func AddUsersEmailListToCoupon (data dictionary.AddUserListToCouponRequest) (string, error) {
	db := database.GetDB()
	fmt.Println("got Here")
	fmt.Println(data)
	for _, userEmail := range data.UsersEmail {
		query := `
			INSERT INTO public.user_coupon(
			user_id, coupon_id, is_used, created_at, updated_at)
			VALUES 
			((
				SELECT id 
				FROM public.user
				WHERE email = $1
				LIMIT 1
			), $2, false, NOW(), null);
		`
		result, err := db.Exec(query,
			userEmail,
			data.CouponId,
		)

		if err != nil {
			fmt.Println("errror ")
			fmt.Println(err)
			return "", err
		}

		affected, err := result.RowsAffected()
		if err != nil {
			return "", err
		}

		if affected == 0 {
			return "", errors.New("no row created")
		}
	} 

	return "success", nil
}


func CreateCoupon (data dictionary.CouponRequest) (*dictionary.Coupon, error) {
	db := database.GetDB()
	createdAt := time.Now()
	query := `
		INSERT INTO public.coupon(
		coupon_name, banner_url, category, type, coupon_percentage, coupon_max_nominal, min_transaction, start_date, expire_date, terms_and_condition, guide, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);
	`
	result, err := db.Exec(query,
		data.CouponName,
		data.BannerUrl,
		data.Category,
		data.Type,
		data.CouponPercentage,
		data.CouponMaxNominal,
		data.MinTransaction,
		data.StartDate,
		data.ExpireDate,
		data.TermsAndCondition,
		data.Guide,
		createdAt,
	)

	if err != nil {
		fmt.Println("errror ")
		fmt.Println(err)
		return nil, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if affected == 0 {
		return nil, errors.New("no row created")
	}

	var couponRes dictionary.Coupon
	id, err := result.LastInsertId()
	couponRes.ID = id
	couponRes.BannerUrl = data.BannerUrl
	couponRes.Category = data.Category
	couponRes.CouponMaxNominal = data.CouponMaxNominal
	couponRes.CouponName = data.CouponName
	couponRes.CouponPercentage = data.CouponPercentage
	couponRes.CreatedAt = createdAt.String()
	couponRes.ExpireDate = data.ExpireDate
	couponRes.Guide = data.Guide
	couponRes.MinTransaction = data.MinTransaction
	couponRes.StartDate = data.StartDate
	couponRes.TermsAndCondition = data.TermsAndCondition
	couponRes.Type = data.Type
	couponRes.UpdatedAt = ""
	return &couponRes, nil
}