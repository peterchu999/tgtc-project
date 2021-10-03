package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/radityaqb/tgtc/backend/database"
	model "github.com/radityaqb/tgtc/backend/dictionary"
)

func UpdateCouponWhenLive(coupon model.Coupon) (interface{}, error) {
	var c model.Coupon
	query := `
	SELECT * 
	FROM public.coupon
	WHERE id = $1`

	db := database.GetDB()
	err := db.QueryRow(query, coupon.ID).Scan(
		&c.ID,
		&c.CouponName,
		&c.BannerUrl,
		&c.Category,
		&c.Type,
		&c.CouponPercentage,
		&c.CouponMaxNominal,
		&c.MinTransaction,
		&c.StartDate,
		&c.ExpireDate,
		&c.TermsAndCondition,
		&c.Guide,
		&c.CreatedAt,
		&c.UpdatedAt,
	)
	if err != nil {
		fmt.Println("coupon with id " + strconv.Itoa(int(coupon.ID)) + " not Found")
		return "coupon with id " + strconv.Itoa(int(coupon.ID)) + " not Found", nil
	}

	startCoupon, err := time.Parse(time.RFC3339, c.StartDate)
	if err != nil {
		fmt.Println(err)
		return "Error parsing start Date", err
	}
	endCoupon, err := time.Parse(time.RFC3339, c.ExpireDate)
	if err != nil {
		fmt.Println(err)
		return "Error parsing expire Date", err
	}
	now := time.Now()

	if now.After(startCoupon) && now.Before(endCoupon) {
		return "Cant update live coupon", nil
	}

	if coupon.CouponName != "" {
		c.CouponName = coupon.CouponName
	}
	if coupon.BannerUrl != "" {
		c.BannerUrl = coupon.BannerUrl
	}
	if coupon.Category != "" {
		c.Category = coupon.Category
	}
	if coupon.Type != "" {
		c.Type = coupon.Type
	}
	if coupon.CouponPercentage != 0 {
		c.CouponPercentage = coupon.CouponPercentage
	}
	if coupon.CouponMaxNominal != 0 {
		c.CouponMaxNominal = coupon.CouponMaxNominal
	}
	if coupon.MinTransaction != 0 {
		c.MinTransaction = coupon.MinTransaction
	}
	if coupon.StartDate != "" {
		c.StartDate = coupon.StartDate
	}
	if coupon.ExpireDate != "" {
		c.ExpireDate = coupon.ExpireDate
	}
	if coupon.TermsAndCondition != "" {
		c.TermsAndCondition = coupon.TermsAndCondition
	}
	if coupon.Guide != "" {
		c.Guide = coupon.Guide
	}
	if coupon.CreatedAt != "" {
		c.CreatedAt = coupon.CreatedAt
	}
	if coupon.UpdatedAt != "" {
		c.UpdatedAt = coupon.UpdatedAt
	}

	updateQuery := `
	UPDATE public.coupon
	SET coupon_name = $1, banner_url = $2, category = $3, coupon_percentage = $4, coupon_max_nominal = $5, min_transaction = $6, start_date = $7, expire_date = $8, terms_and_condition = $8, guide = $9, created_at = $10, updated_at = $11
	WHERE id = $12
	`
	startDate, _ := time.Parse(time.RFC3339, c.StartDate)
	expireDate, _ := time.Parse(time.RFC3339, c.ExpireDate)
	createdAt, _ := time.Parse(time.RFC3339, c.CreatedAt)
	updatedAt, _ := time.Parse(time.RFC3339, c.UpdatedAt)
	res, err := db.Exec(
		updateQuery,
		c.CouponName,
		c.BannerUrl,
		c.Category,
		c.Type,
		c.CouponPercentage,
		c.CouponMaxNominal,
		c.MinTransaction,
		startDate,
		expireDate,
		c.TermsAndCondition,
		c.Guide,
		createdAt,
		updatedAt,
		c.ID,
	)
	if err != nil {
		fmt.Println(err)
		return "Failed update coupon", err
	}

	count, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return "Failed update coupon", err
	}

	if count == 1 {
		fmt.Println("Success update coupon")
		return "Success update coupon", nil
	} else {
		fmt.Println(err)
		return "Failed update coupon", err
	}
}

func GetCouponWithUserEmail(email string) ([]model.Coupon, error) {
	db := database.GetDB()
	query := `
	SELECT 
		public.coupon.*
	FROM 
		public.coupon JOIN
		(SELECT * 
		FROM public.user JOIN public.user_coupon ON public.user.id = public.user_coupon.user_id
		WHERE public.user.email = $1) AS u
		ON public.coupon.id = u.coupon_id
	`
	rows, err := db.Query(query, email)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var coupons []model.Coupon
	for rows.Next() {
		var c model.Coupon
		err = rows.Scan(
			&c.ID,
			&c.CouponName,
			&c.BannerUrl,
			&c.Category,
			&c.Type,
			&c.CouponPercentage,
			&c.CouponMaxNominal,
			&c.MinTransaction,
			&c.StartDate,
			&c.ExpireDate,
			&c.TermsAndCondition,
			&c.Guide,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		fmt.Println(c)
		coupons = append(coupons, c)
	}

	return coupons, nil
}
