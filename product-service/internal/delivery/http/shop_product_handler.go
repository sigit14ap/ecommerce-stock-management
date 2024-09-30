package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/sigit14ap/product-service/helpers"
	"github.com/sigit14ap/product-service/internal/delivery/dto"
	"github.com/sigit14ap/product-service/internal/domain"
	"github.com/sigit14ap/product-service/internal/usecase"
)

type ProductHandler struct {
	productUsecase usecase.ProductUsecase
}

var validate *validator.Validate

func NewProductHandler(productUsecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUsecase: productUsecase}
}

func (handler *ProductHandler) GetAllByShopID(context *gin.Context) {
	shopIDValue, exists := context.Get("shopID")

	if !exists {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "Shop ID not found")
		return
	}

	shopID, ok := shopIDValue.(uint64)
	if !ok {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "invalid Shop ID")
		return
	}

	products, err := handler.productUsecase.GetAllByShopID(shopID)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(context, products)
}

func (handler *ProductHandler) GetByIDAndShopID(context *gin.Context) {
	shopIDValue, exists := context.Get("shopID")

	if !exists {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "Shop ID not found")
		return
	}

	shopID, ok := shopIDValue.(uint64)
	if !ok {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "invalid Shop ID")
		return
	}

	id := context.Param("id")

	ProductID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "invalid Product ID")
		return
	}
	product, err := handler.productUsecase.GetByIDAndShopID(ProductID, shopID)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(context, product)
}

func (handler *ProductHandler) Create(context *gin.Context) {

	validate = validator.New()
	var productRequest dto.ProductRequest
	err := context.BindJSON(&productRequest)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = validate.Struct(productRequest)
	if err != nil {
		helpers.ErrorValidationResponse(context, err)
		return
	}

	shopIDValue, exists := context.Get("shopID")

	if !exists {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "Shop ID not found")
		return
	}

	shopID, ok := shopIDValue.(uint64)
	if !ok {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "invalid Shop ID")
		return
	}

	product := &domain.Product{
		Name:   productRequest.Name,
		Price:  productRequest.Price,
		ShopID: shopID,
	}

	if err := handler.productUsecase.Create(product); err != nil {
		helpers.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(context, product)
}

func (handler *ProductHandler) Update(context *gin.Context) {

	validate = validator.New()
	var productRequest dto.ProductRequest
	err := context.BindJSON(&productRequest)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	err = validate.Struct(productRequest)
	if err != nil {
		helpers.ErrorValidationResponse(context, err)
		return
	}

	shopIDValue, exists := context.Get("shopID")

	if !exists {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "Shop ID not found")
		return
	}

	shopID, ok := shopIDValue.(uint64)
	if !ok {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "invalid Shop ID")
		return
	}

	id := context.Param("id")

	ProductID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "invalid Product ID")
		return
	}

	product := &domain.Product{
		ID:     ProductID,
		Name:   productRequest.Name,
		Price:  productRequest.Price,
		ShopID: shopID,
	}

	if err := handler.productUsecase.Update(product); err != nil {
		helpers.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(context, product)
}

func (handler *ProductHandler) Delete(context *gin.Context) {
	shopIDValue, exists := context.Get("shopID")

	if !exists {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "Shop ID not found")
		return
	}

	shopID, ok := shopIDValue.(uint64)
	if !ok {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "invalid Shop ID")
		return
	}

	id := context.Param("id")

	ProductID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		helpers.ErrorResponse(context, http.StatusUnauthorized, "invalid Product ID")
		return
	}

	if err := handler.productUsecase.Delete(ProductID, shopID); err != nil {
		helpers.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(context, nil)
}
