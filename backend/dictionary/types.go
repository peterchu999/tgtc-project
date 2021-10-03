package dictionary


type CouponRequest struct {
	CouponName        string `json:"coupon_name"`
	BannerUrl         string `json:"banner_url"`
	Category          string `json:"category"`
	Type              string `json:"type"`
	CouponPercentage  float64 `json:"coupon_percentage"`
	CouponMaxNominal  float64 `json:"coupon_max_nominal"`
	MinTransaction    float64 `json:"min_transaction"`
	StartDate         string `json:"start_date"`
	ExpireDate        string `json:"expire_date"`
	TermsAndCondition string `json:"terms_and_condition"`
	Guide             string `json:"guide"`
}

type AddUserListToCouponRequest struct {
	CouponId int64 `json:coupon_id`
	UsersEmail [2]string `json:users_email`
}
type Coupon struct {
	ID                int64  `json:"id"`
	CouponName        string `json:"coupon_name"`
	BannerUrl         string `json:"banner_url"`
	Category          string `json:"category"`
	Type              string `json:"type"`
	CouponPercentage  float64 `json:"coupon_percentage"`
	CouponMaxNominal  float64 `json:"coupon_max_nominal"`
	MinTransaction    float64 `json:"min_transaction"`
	StartDate         string `json:"start_date"`
	ExpireDate        string `json:"expire_date"`
	TermsAndCondition string `json:"terms_and_condition"`
	Guide             string `json:"guide"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
}


type Product struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	ShopName     string  `json:"shop_name"`
	ProductPrice float64 `json:"product_price"`
	ImageURL     string  `json:"image_url"`
}

type User struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	BirthDate   string `json:"birth_date"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Membership  string `json:"membership"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type APIResponse struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error_message"`
}

