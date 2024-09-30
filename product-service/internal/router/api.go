package router

import (
	delivery "github.com/sigit14ap/product-service/internal/delivery/http"
	"github.com/sigit14ap/product-service/internal/middleware"
	"github.com/sigit14ap/product-service/internal/services"

	"github.com/gin-gonic/gin"
)

func NewRouter(productHandler *delivery.ProductHandler, userProductHandler *delivery.UserProductHandler, shopClient *services.ShopClient) *gin.Engine {
	router := gin.New()
	v1 := router.Group("/api/v1")
	v1.Use(middleware.ServiceMiddleware())

	shop := v1.Group("shop")
	shop.Use(middleware.ShopMiddleware(shopClient))
	{
		shop.GET("/products", productHandler.GetAllByShopID)
		shop.POST("/products", productHandler.Create)
		shop.GET("/products/:id", productHandler.GetByIDAndShopID)
		shop.PUT("/products/:id", productHandler.Update)
		shop.DELETE("/products/:id", productHandler.Delete)
	}

	products := v1.Group("products")
	{
		products.GET("/", userProductHandler.GetAllProductsWithStock)
	}

	return router
}
