package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/cmd"
	config "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	adminlogs "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/admin_logs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/admins"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/auth"
	cartitems "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/cart_items"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/carts"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/categories"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/image"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/orders"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/payments"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/products"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/reports"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/reviews"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/shipments"
	systembank "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/system_bank"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/users"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/wishlist"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/middlewares"

	"github.com/spf13/cobra"
)

func main() {
	config.Database()
	if err := command(); err != nil {
		log.Fatalf("Error runing command :%v", err)
	}
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowMethods:           []string{"*"},
		AllowHeaders:           []string{"*"},
		AllowCredentials:       true,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		AllowWebSockets:        true,
		AllowFiles:             false,
	}))

	md := middlewares.AuthMiddleware()

	// User
	r.POST("/user/create", users.CreateUser)
	r.GET("/user", md, users.UserList)
	r.GET("/user/:id", md, users.GetUserByID)
	r.DELETE("/user/:id", md, users.DeleteUser)
	r.PATCH("/user/:id", md, users.UpdateUser)

	// Get Info
	r.GET("/user/info", md, users.GetInfoUser)
	r.GET("/admin/info", md, admins.GetInfoAdmin)

	// Product
	r.POST("/product/create", md, products.CreateProduct)
	r.GET("/product/:id", products.GetProductByID)
	r.GET("/product", products.ProductList)
	r.DELETE("/product/:id", md, products.DeleteProduct)
	r.PATCH("/product/:id", md, products.UpdateProduct)

	//auth
	r.POST("/auth/login", auth.LoginUser)
	r.POST("/auth/login/admin", auth.LoginAdmin)

	// Order
	r.POST("/order/create", md, orders.CreateOrder)
	r.GET("/order/:id", md, orders.GetOrderByID)
	r.GET("/order", md, orders.OrderList)
	// r.GET("/order/user", md, orders.OrderUserList)
	//r.DELETE("/order/:id", orders.DeleteOrder)
	r.PATCH("/order/:id", md, orders.UpdateOrder)

	// Order User
	r.GET("/order/pending", md, orders.OrderUserPendingList)
	r.GET("/order/paid", md, orders.OrderUserPaidList)
	r.GET("/order/prepare", md, orders.OrderUserPrepareList)
	r.GET("/order/ship", md, orders.OrderUserShipList)
	r.GET("/order/success", md, orders.OrderUserSuccessList)
	r.GET("/order/failed", md, orders.OrderUserFailedList)
	r.GET("/order/cancelled", md, orders.OrderUserCancelledList)
	r.GET("/order/history", md, orders.OrderUserHistoryList)
	r.PATCH("/order/ship/:id", md, orders.UpdateShipOrder)

	// Wishlist
	r.POST("/wish/create", md, wishlist.CreateWishlist)
	// r.GET("/wish", md, wishlist.GetWishlistByID)
	r.GET("/wish", md, wishlist.Wishlist)
	r.DELETE("/wish/:id", md, wishlist.DeleteWishlists)
	r.PATCH("/wish/update", md, wishlist.UpdateWishlists)
	r.GET("/wish/status", md, wishlist.GetWishlistStatus)

	// Cart
	r.POST("/cart/create", carts.AddCart)
	r.GET("/cart", md, carts.GetCartByID)
	// r.GET("/cart", carts.CartList)
	r.DELETE("/cart/:id", md, carts.DeleteCart)
	r.PATCH("/cart/:id", md, carts.UpdateCart)

	// CartItem
	r.POST("/cartitem/create", md, cartitems.CreateCartItem)
	r.GET("/cartitem", md, cartitems.CartItemList)
	r.DELETE("/cartitem", md, cartitems.DeleteCartItem)
	r.PATCH("/cartitem/:id", md, cartitems.UpdateCartItem)

	// r.GET("/cartitem/:id", cartitems.GetCartItemByID)
	// r.GET("/cartitem", cartitems.CartItemList)
	// r.DELETE("/cartitem", cartitems.DeleteCartItem)
	// r.PATCH("/cartitem/:id", cartitems.UpdateCartItem)

	// Review
	r.POST("/review/create", md, reviews.CreateReview)
	r.GET("/review/user", md, reviews.GetReviewByID)
	r.GET("/review", md, reviews.ReviewList)
	r.DELETE("/review/:id", md, reviews.DeleteReview)
	r.PATCH("/review/:id", md, reviews.UpdateReview)

	// Admin
	r.POST("/admin/create", md, admins.CreateAdmin)
	r.GET("/admin/:id", md, admins.GetAdminByID)
	r.GET("/admin", md, admins.AdmintList)
	r.DELETE("/admin/:id", md, admins.DeleteAdmin)
	r.PATCH("/admin/:id", md, admins.UpdateAdmin)

	// Admin_log
	r.GET("/adminlog", md, adminlogs.AdminLogList)

	// System bank
	r.POST("/system/create", md, systembank.CreateSystembank)
	r.GET("/system/:id", md, systembank.GetSystemBankByID)
	r.GET("/system", md, systembank.SystemBankList)
	r.DELETE("/system/:id", md, systembank.DeleteSystemBank)
	r.PATCH("/system/:id", md, systembank.UpdateSystemBank)

	// Category
	r.POST("/category/create", md, categories.CreateCategory)
	r.GET("/category/:id", categories.GetCategoryByID)
	r.GET("/category", categories.CategoryList)
	r.DELETE("/category/:id", md, categories.DeleteCategory)
	r.PATCH("/category/:id", md, categories.UpdateCategory)

	// Shipment
	r.POST("/shipment/create", shipments.CreateShipment)
	r.GET("/shipment/:id", shipments.GetShipmentByID)
	r.GET("/shipment", md, shipments.ShipmenttList)
	r.DELETE("/shipment/:id", shipments.DeleteShipment)
	r.PATCH("/shipment/:id", shipments.UpdateShipment)

	// Payment
	r.POST("/payment/create", payments.CreatePayment)
	r.GET("/payment/:id", payments.GetPaymentByID)
	r.GET("/payment", payments.PaymentList)
	r.DELETE("/payment/:id", payments.DeletePayment)
	r.PATCH("/payment/:id", md, payments.UpdatePayment)

	// banner
	r.GET("/banner", md, image.BannerList)
	r.POST("/banner/create", md, image.CreateImageBanner)
	r.DELETE("/banner/:id", md, image.DeleteImageBanner)

	// Dashboard
	r.GET("/dashboard", reports.Dashboard)
	r.GET("/report", reports.Report)
	r.GET("/dashboard/category", reports.DashboardTotalByCategory)

	r.Run()

}

func command() error {
	cmda := &cobra.Command{
		Use:  "app",
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	cmda.AddCommand(cmd.Migrate())

	return cmda.Execute()
}
