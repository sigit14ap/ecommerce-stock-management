package usecase

import (
	"github.com/sigit14ap/product-service/internal/delivery/dto"
	"github.com/sigit14ap/product-service/internal/domain"
	repository "github.com/sigit14ap/product-service/internal/repository/mysql"
)

type ProductUsecase interface {
	GetAllByShopID(shopID uint64) ([]domain.Product, error)
	GetByIDAndShopID(id, shopID uint64) (*domain.Product, error)
	Create(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(id, shopID uint64) error
	GetAllProductsWithStock() ([]dto.ProductResponse, error)
}

type productUsecase struct {
	productRepository repository.ProductRepository
}

func NewProductUsecase(productRepository repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepository: productRepository,
	}
}

func (uc *productUsecase) GetAllByShopID(shopID uint64) ([]domain.Product, error) {
	return uc.productRepository.GetAllByShopID(shopID)
}

func (uc *productUsecase) GetByIDAndShopID(id, shopID uint64) (*domain.Product, error) {
	return uc.productRepository.GetByIDAndShopID(id, shopID)
}

func (uc *productUsecase) Create(product *domain.Product) error {
	return uc.productRepository.Create(product)
}

func (uc *productUsecase) Update(product *domain.Product) error {
	return uc.productRepository.Update(product)
}

func (uc *productUsecase) Delete(id, shopID uint64) error {
	return uc.productRepository.Delete(id, shopID)
}

func (uc *productUsecase) GetAllProductsWithStock() ([]dto.ProductResponse, error) {
	return uc.productRepository.GetAllProductsWithStock()
}
