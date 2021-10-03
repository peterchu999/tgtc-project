package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/peterchu999/tgtc/backend/dictionary"
	"github.com/peterchu999/tgtc/backend/service"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

// func AddProduct(w http.ResponseWriter, r *http.Request) {

// 	var p dictionary.Product

// 	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
// 		http.Error(w, "bad request", 400)
// 		return
// 	}

// 	p = product.AddProduct(context.Background(), p)
// 	fmt.Fprintf(w, fmt.Sprint("success, id product : ", p.ID))
// }

// func GetProduct(w http.ResponseWriter, r *http.Request) {
// 	idstring := r.URL.Query().Get("id")

// 	idInt64, err := strconv.ParseInt(idstring, 10, 64)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	p, err := product.GetProduct(context.Background(), idInt64)
// 	if err != nil {
// 		// log.Fatal(err)
// 		fmt.Fprintf(w, err.Error())
// 	}

// 	val, err := json.Marshal(p)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Fprintf(w, string(val))
// }

// func DeleteProduct(w http.ResponseWriter, r *http.Request) {
// 	var p dictionary.Product

// 	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
// 		http.Error(w, "bad request", 400)
// 		return
// 	}

// 	product.DeleteProduct(context.Background(), p.ID)
// 	fmt.Fprintf(w, "success")
// }

// func UpdateProduct(w http.ResponseWriter, r *http.Request) {
// 	var p dictionary.Product

// 	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
// 		http.Error(w, "bad request", 400)
// 		return
// 	}

// 	product.UpdateProduct(context.Background(), p)

// 	fmt.Fprintf(w, "success")
// }


func AddCoupon(w http.ResponseWriter, r *http.Request) {

	var c dictionary.CouponRequest
	fmt.Print(r.Body)
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	coupon, _ := service.CreateCoupon(c)
	b, err := json.Marshal(coupon)
    if err != nil {
        fmt.Println(err)
        return
    }
	fmt.Fprintf(w, string(b))
}

func AddUsersToCouponByEmail(w http.ResponseWriter, r *http.Request) {

	var addUserListToCouponRequest dictionary.AddUserListToCouponRequest
	fmt.Print(r.Body)
	if err := json.NewDecoder(r.Body).Decode(&addUserListToCouponRequest); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	fmt.Print(addUserListToCouponRequest)
	addUserListToCouponRequest.CouponId = 1
	addUserListToCouponRequest.UsersEmail[0] = "test@test.com"
	coupon, _ := service.AddUsersEmailListToCoupon(addUserListToCouponRequest)
	b, err := json.Marshal(coupon)
    if err != nil {
        fmt.Println(err)
        return
    }
	fmt.Fprintf(w, string(b))
}

func UpdateCoupon(w http.ResponseWriter, r *http.Request) {

	var c dictionary.Coupon
	fmt.Print(r.Body)
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	coupon, _ := service.UpdateCouponWhenNotLive(c)
	b, err := json.Marshal(coupon)
    if err != nil {
        fmt.Println(err)
        return
    }
	fmt.Fprintf(w, string(b))
}

func GetUsersToCouponByEmail(w http.ResponseWriter, r *http.Request) {

	var getCouponToUserEmailRequest string
	fmt.Print(r.Body)
	if err := json.NewDecoder(r.Body).Decode(&getCouponToUserEmailRequest); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	fmt.Print(getCouponToUserEmailRequest)
	// addUserListToCouponRequest.CouponId = 1
	// addUserListToCouponRequest.UsersEmail[0] = "test@test.com"
	coupon, _ := service.GetCouponWithUserEmail(getCouponToUserEmailRequest)
	b, err := json.Marshal(coupon)
    if err != nil {
        fmt.Println(err)
        return
    }
	fmt.Fprintf(w, string(b))
}

// func GetProduct(w http.ResponseWriter, r *http.Request) {
// 	idstring := r.URL.Query().Get("id")

// 	idInt64, err := strconv.ParseInt(idstring, 10, 64)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	p, err := product.GetProduct(context.Background(), idInt64)
// 	if err != nil {
// 		// log.Fatal(err)
// 		fmt.Fprintf(w, err.Error())
// 	}

// 	val, err := json.Marshal(p)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Fprintf(w, string(val))
// }

// func DeleteProduct(w http.ResponseWriter, r *http.Request) {
// 	var p dictionary.Product

// 	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
// 		http.Error(w, "bad request", 400)
// 		return
// 	}

// 	product.DeleteProduct(context.Background(), p.ID)
// 	fmt.Fprintf(w, "success")
// }

// func UpdateProduct(w http.ResponseWriter, r *http.Request) {
// 	var p dictionary.Product

// 	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
// 		http.Error(w, "bad request", 400)
// 		return
// 	}

// 	product.UpdateProduct(context.Background(), p)

// 	fmt.Fprintf(w, "success")
// }
