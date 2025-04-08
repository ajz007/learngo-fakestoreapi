package client

import "github.com/ajz007/learngo-fakestoreapi/internal/model"

type FakeStoreProductDTO struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Image       string    `json:"image"`
	Rating      RatingDTO `json:"rating"`
}

type RatingDTO struct {
	Rate  float64 `json:"rate"`
	Count int     `json:"count"`
}

func (dto FakeStoreProductDTO) ToProduct() model.Product {
	return model.Product{
		ID:          dto.ID,
		Title:       dto.Title,
		Price:       dto.Price,
		Description: dto.Description,
		Category:    dto.Category,
		Image:       dto.Image,
		Rating:      dto.Rating.toRating(),
	}
}

func (dto RatingDTO) toRating() model.Rating {
	return model.Rating{
		Rate:  dto.Rate,
		Count: dto.Count,
	}
}
