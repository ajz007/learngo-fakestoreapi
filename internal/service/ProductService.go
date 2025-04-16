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

func NewProductService(client client.FakeStoreClient) ProductService {
	return ProductServiceImpl{
		fakestoreClient: client,
	}
}

func (ps ProductServiceImpl) GetProducts() ([]model.Product, error) {
	products, err := ps.fakestoreClient.GetProducts()

	if err != nil {
		return nil, err
	}

	return products, nil
}
