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
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/orders"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/payments"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/products"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/reviews"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/shipments"
	systembank "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/system_bank"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/users"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/wishlist"

	// "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/middlewares"
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

	// md := middlewares.AuthMiddleware()

	// User
	r.POST("/user/create", users.CreateUser)
	r.GET("/user/:id", users.GetUserByID)
	r.DELETE("/user/:id", users.DeleteUser)
	r.PATCH("/user/:id", users.UpdateUser)

	// Product
	r.POST("/product/create", products.CreateProduct)
	r.GET("/product/:id", products.GetProductByID)
	r.GET("/product", products.ProductList)
	r.DELETE("/product/:id", products.DeleteProduct)
	r.PATCH("/product/:id", products.UpdateProduct)

	//auth
	r.POST("/auth/login", auth.LoginUser)
	r.POST("/auth/login/admin", auth.LoginAdmin)

	// Order
	r.POST("/order/create", orders.CreateOrder)
	r.GET("/order/:id", orders.GetOrderByID)
	r.GET("/order", orders.OrderList)
	r.DELETE("/order/:id", orders.DeleteOrder)
	r.PATCH("/order/:id", orders.UpdateOrder)

	// Wishlist
	r.POST("/wish/create", wishlist.CreateWishlist)
	r.GET("/wish/:id", wishlist.GetWishlistByID)
	r.GET("/wish", wishlist.Wishlist)
	r.DELETE("/wish/:id", wishlist.DeleteWishlists)
	r.PATCH("/wish/:id", wishlist.UpdateWishlists)

	// Cart
	r.POST("/cart/create", carts.AddCart)
	r.GET("/cart/:id", carts.GetCartByID)
	r.GET("/cart", carts.CartList)
	r.DELETE("/cart/:id", carts.DeleteCart)
	r.PATCH("/cart/:id", carts.UpdateCart)

	// CartItem
	r.POST("/cartitem/create", cartitems.CreateCartItem)
	r.GET("/cartitem/:id", cartitems.GetCartItemByID)
	r.GET("/cartitem", cartitems.CartItemList)
	r.DELETE("/cartitem/:id", cartitems.DeleteCartItem)
	r.PATCH("/cartitem/:id", cartitems.UpdateCartItem)

	// Review
	r.POST("/review/create", reviews.CreateReview)
	r.GET("/review/:id", reviews.GetReviewByID)
	r.GET("/review", reviews.ReviewList)
	r.DELETE("/review/:id", reviews.DeleteReview)
	r.PATCH("/review/:id", reviews.UpdateReview)

	// Admin
	r.POST("/admin/create", admins.CreateAdmin)
	r.GET("/admin/:id", admins.GetAdminByID)
	r.GET("/admin", admins.AdmintList)
	r.DELETE("/admin/:id", admins.DeleteAdmin)
	r.PATCH("/admin/:id", admins.UpdateAdmin)

	// Admin_log
	r.GET("/adminlog", adminlogs.AdminLogList)

	// System bank
	r.POST("/system/create", systembank.CreateSystembank)
	r.GET("/system/:id", systembank.GetSystemBankByID)
	r.GET("/system", systembank.SystemBankList)
	r.DELETE("/system/:id", systembank.DeleteSystemBank)
	r.PATCH("/system/:id", systembank.UpdateSystemBank)

	// Category
	r.POST("/category/create", categories.CreateCategory)
	r.GET("/category/:id", categories.GetCategoryByID)
	r.GET("/category", categories.CategoryList)
	r.DELETE("/category/:id", categories.DeleteCeategory)
	r.PATCH("/category/:id", categories.UpdateCategory)

	// Shipment
	r.POST("/shipment/create", shipments.CreateShipment)
	r.GET("/shipment/:id", shipments.GetShipmentByID)
	r.GET("/shipment", shipments.ShipmenttList)
	r.DELETE("/shipment/:id", shipments.DeleteShipment)
	r.PATCH("/shipment/:id", shipments.UpdateShipment)

	// Payment
	r.POST("/payment/create", payments.CreatePayment)
	r.GET("/payment/:id", payments.GetPaymentByID)
	r.GET("/payment", payments.PaymentList)
	r.DELETE("/payment/:id", payments.DeletePayment)
	r.PATCH("/payment/:id", payments.UpdatePayment)

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
