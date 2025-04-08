package service

import (
	"github.com/ajz007/learngo-fakestoreapi/internal/client"
	"github.com/ajz007/learngo-fakestoreapi/internal/model"
)

type ProductService interface {
	GetProducts() ([]model.Product, error)
}

type ProductServiceImpl struct {
	fakestoreClient client.FakeStoreClient
}

func NewProductService() ProductService {
	return ProductServiceImpl{
		fakestoreClient: client.NewFakeStoreClient(),
	}
}

func (ps ProductServiceImpl) GetProducts() ([]model.Product, error) {
	productsDto, err := ps.fakestoreClient.GetProducts()

	if err != nil {
		return nil, err
	}

	var products []model.Product

	for _, productDto := range productsDto {
		products = append(products, productDto.ToProduct())
	}

	return products, nil
}
