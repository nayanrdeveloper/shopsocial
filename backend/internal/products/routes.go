package products

import (
	"shopsocial-backend/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.RouterGroup, handler *ProductHandler) {
	productRoutes := router.Group("/products")
	productRoutes.Use(middleware.AuthMiddleware())
	{
		productRoutes.POST("/", handler.CreateProduct)
		productRoutes.GET("/:id", handler.GetProductByID)
		productRoutes.PUT("/:id", handler.UpdateProduct)
		productRoutes.DELETE("/:id", handler.DeleteProduct)
	}
}
