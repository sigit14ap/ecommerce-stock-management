package repository

import (
	"github.com/sigit14ap/product-service/internal/delivery/dto"
	"github.com/sigit14ap/product-service/internal/domain"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllByShopID(shopID uint64) ([]domain.Product, error)
	GetByIDAndShopID(id, shopID uint64) (*domain.Product, error)
	Create(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(id, shopID uint64) error
	GetAllProductsWithStock() ([]dto.ProductResponse, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (repository *productRepository) GetAllByShopID(shopID uint64) ([]domain.Product, error) {
	var products []domain.Product
	if err := repository.db.Where("shop_id = ?", shopID).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (repository *productRepository) GetByIDAndShopID(id, shopID uint64) (*domain.Product, error) {
	var product domain.Product
	if err := repository.db.Where("id = ? AND shop_id = ?", id, shopID).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (repository *productRepository) Create(product *domain.Product) error {
	return repository.db.Create(product).Error
}

func (repository *productRepository) Update(product *domain.Product) error {
	return repository.db.Omit("CreatedAt").Save(product).Error
}

func (repository *productRepository) Delete(id, shopID uint64) error {
	return repository.db.Where("id = ? AND shop_id = ?", id, shopID).Delete(&domain.Product{}).Error
}

func (repository *productRepository) GetAllProductsWithStock() ([]dto.ProductResponse, error) {
	var products []dto.ProductResponse

	err := repository.db.
		Model(&domain.Product{}).
		Select("products.*, IFNULL(SUM(stocks.quantity), 0) as total_stock").
		Joins("LEFT JOIN stocks ON stocks.product_id = products.id").
		Group("products.id").
		Scan(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}
