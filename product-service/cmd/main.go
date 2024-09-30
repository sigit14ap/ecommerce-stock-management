package main

import (
	"fmt"
	"os"

	"github.com/sigit14ap/product-service/config"
	"github.com/sigit14ap/product-service/helpers"
	delivery "github.com/sigit14ap/product-service/internal/delivery/http"
	"github.com/sigit14ap/product-service/internal/domain"
	repository "github.com/sigit14ap/product-service/internal/repository/mysql"
	"github.com/sigit14ap/product-service/internal/router"
	"github.com/sigit14ap/product-service/internal/services"
	"github.com/sigit14ap/product-service/internal/usecase"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()

	log := helpers.InitializeLogs()

	if log == nil {
		log.Fatal("Logger failed to started")
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DatabaseUser,
		cfg.DatabasePassword,
		cfg.DatabaseHost,
		cfg.DatabasePort,
		cfg.DatabaseName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&domain.Product{})
	if err != nil {
		log.Fatalf("failed to auto-migrate Product model: %v", err)
	}

	shopClient := services.NewShopClient(cfg.ShopServiceUrl, cfg.AppSecret)

	productRepo := repository.NewProductRepository(db)
	productService := usecase.NewProductUsecase(productRepo)
	productHandler := delivery.NewProductHandler(productService)
	userProductHandler := delivery.NewUserProductHandler(productService)

	router := router.NewRouter(productHandler, userProductHandler, shopClient)

	log.Info(router.Run(":" + os.Getenv("APP_PORT")))
}
