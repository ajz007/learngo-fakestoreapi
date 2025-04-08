package handler

import (
	"net/http"

	"github.com/ajz007/learngo-fakestoreapi/internal/service"
	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	GetProducts(c *gin.Context)
}

type ProductHandlerImpl struct {
	prodService service.ProductService
}

func NewProductHandler() ProductHandler {
	return &ProductHandlerImpl{
		prodService: service.NewProductService(),
	}
}

func (ph *ProductHandlerImpl) GetProducts(c *gin.Context) {
	products, err := ph.prodService.GetProducts()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, products)
	}
}
