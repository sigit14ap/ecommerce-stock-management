package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sigit14ap/product-service/helpers"
	"github.com/sigit14ap/product-service/internal/usecase"
)

type UserProductHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewUserProductHandler(productUsecase usecase.ProductUsecase) *UserProductHandler {
	return &UserProductHandler{productUsecase: productUsecase}
}

func (handler *UserProductHandler) GetAllProductsWithStock(context *gin.Context) {
	products, err := handler.productUsecase.GetAllProductsWithStock()
	if err != nil {
		helpers.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	helpers.SuccessResponse(context, products)
}
