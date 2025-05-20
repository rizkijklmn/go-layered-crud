package main

import (
	"go-crud/config"
	"go-crud/product"
	"go-crud/routes"
	"go-crud/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDB()

	// Auto migrate
	// config.DB.Migrator().DropTable(&product.Product{}) // Drop tabel
	config.DB.AutoMigrate(&user.User{}, &product.Product{})

	routes.SetupRoutes(r)

	r.Run(":3000") // Jalankan server di port 3000
}
