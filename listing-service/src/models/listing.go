package models

import "time"

type Listing struct {
	ID          uint
	UserId      int16
	Price       int64
	ListingType string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
