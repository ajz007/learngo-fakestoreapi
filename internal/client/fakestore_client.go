package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ajz007/learngo-fakestoreapi/internal/model"
)

type FakeStoreClient interface {
	GetProducts() ([]model.Product, error)
}

type fakeStoreClientImpl struct {
	baseUrl string
	client  *http.Client
}

func NewFakeStoreClient() FakeStoreClient {
	return &fakeStoreClientImpl{
		baseUrl: "https://fakestoreapi.com",
		client: &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       5 * time.Second,
		},
	}
}

func (fc *fakeStoreClientImpl) GetProducts() ([]model.Product, error) {
	resp, err := fc.client.Get(fmt.Sprintf(fc.baseUrl + "/products")) //TODO: check if contatenation is ok in golang

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch products from fakestoreapi")
	}

	var productsDTO []FakeStoreProductDTO
	if err := json.NewDecoder(resp.Body).Decode(&productsDTO); err != nil {
		return nil, err
	}

	var products []model.Product

	for _, productDto := range productsDTO {
		products = append(products, productDto.ToProduct())
	}

	return products, nil

}
