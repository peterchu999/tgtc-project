package dictionary

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
