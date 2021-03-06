package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/peterchu999/tgtc/backend/database"
	"github.com/peterchu999/tgtc/backend/handlers"
	"github.com/peterchu999/tgtc/backend/server"
)

func main() {

	// Init database connection
	database.InitDB()

	// Init serve HTTP
	router := mux.NewRouter()

	// routes http
	router.HandleFunc("/ping", handlers.Ping).Methods(http.MethodGet)

	// construct your own API endpoints
	// endpoint : /add-product
	// endpoint : /get-product?id=
	// endpoint : /update-product
	// endpoint : /delete-product
	// router.HandleFunc("/add-product", handlers.AddProduct).Methods(http.MethodPost)
	// router.HandleFunc("/get-product", handlers.GetProduct).Methods(http.MethodGet)
	// router.HandleFunc("/update-product", handlers.UpdateProduct).Methods(http.MethodPatch)
	// router.HandleFunc("/delete-product", handlers.DeleteProduct).Methods(http.MethodDelete)
	router.HandleFunc("/add-coupon", handlers.AddCoupon).Methods(http.MethodPost)
	router.HandleFunc("/add-user-to-coupon-list", handlers.AddUsersToCouponByEmail).Methods(http.MethodPost)
	router.HandleFunc("/update-coupon", handlers.UpdateCoupon).Methods(http.MethodPost)
	router.HandleFunc("/get-by-user", handlers.GetUsersToCouponByEmail).Methods(http.MethodPost)

	serverConfig := server.Config{
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
		Port:         8000,
	}
	server.Serve(serverConfig, router)
}
