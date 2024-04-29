package dtos

type CreateListingDto struct {
	UserId      int16  `json:"user_id"`
	ListingType string `json:"listing_type"`
	Price       int64  `json:"price"`
}
