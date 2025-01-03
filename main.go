package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/cmd"
	config "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/configs"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/admins"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/auth"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/categories"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/products"
	systembank "github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/system_bank"
	"github.com/komkemkku/komkemkku/Back-end_Grit-Electronic/controller/users"

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

	// Product
	r.POST("/product/create", products.CreateProduct)
	r.GET("/product/:id", products.GetProductByID)
	r.GET("/product", products.ProductList)
	r.DELETE("/product/:id", products.DeleteProduct)
	r.PATCH("/product/:id", products.UpdateProduct)

	//auth
	r.POST("/auth/login", auth.LoginUser)
	// r.POST("/auth/login/admin", auth.LoginAdmin)

	// Order

	// Admin
	r.POST("/admin/create", admins.CreateAdmin)
	r.GET("/admin/:id", admins.GetAdminByID)
	r.GET("/admin", admins.AdmintList)
	r.DELETE("/admin/:id", admins.DeleteAdmin)
	r.PATCH("/admin/:id", admins.UpdateAdmin)

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
