package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type FakeStoreClient interface {
	GetProducts() ([]FakeStoreProductDTO, error)
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

func (fc *fakeStoreClientImpl) GetProducts() ([]FakeStoreProductDTO, error) {
	resp, err := fc.client.Get(fmt.Sprintf(fc.baseUrl + "/products")) //TODO: check if contatenation is ok in golang

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch products from fakestoreapi")
	}

	var products []FakeStoreProductDTO
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, err
	}

	return products, nil

}
