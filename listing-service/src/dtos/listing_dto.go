package dtos

import "rent/listing-service/src/models"

type CreateListingDto struct {
	UserId      int16  `json:"user_id" binding:"required"`
	ListingType string `json:"listing_type" binding:"required"`
	Price       int64  `json:"price" binding:"required"`
}

type ListingResponseDto struct {
	Result  bool           `json:"result"`
	Listing models.Listing `json:"listing"`
}

type ListingsResponseDto struct {
	Result   bool             `json:"result"`
	Listings []models.Listing `json:"listings"`
}
