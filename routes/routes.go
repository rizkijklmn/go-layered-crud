package routes

import (
	"go-crud/product"
	"go-crud/user"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		userRoute := api.Group("/users")
		{
			userRoute.POST("", user.CreateUserHandler)
			userRoute.GET("", user.GetUsersHandler)
			userRoute.GET("/:id", user.GetUserByIDHandler)
			userRoute.PUT("/:id", user.UpdateUserHandler)
			userRoute.DELETE("/:id", user.DeleteUserHandler)
		}

		userRoute = api.Group("/products")
		{
			userRoute.POST("", product.CreateProductHandler)
			userRoute.GET("", product.GetProductsHandler)
			userRoute.GET("/:id", product.GetProductByIDHandler)
			userRoute.PUT("/:id", product.UpdateProductHandler)
			userRoute.DELETE("/:id", product.DeleteProductHandler)
		}
	}
}
