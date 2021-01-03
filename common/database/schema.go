package database

import (
	"encoding/json"
)

type (
	// HotelOffers is a collection of offers
	HotelOffers struct {
		Offers []Offer `json:"offers"`
	}
	// Offer contains the hotel, room and rate plan data
	Offer struct {
		Hotel    json.RawMessage `json:"hotel"`
		Room     json.RawMessage `json:"room"`
		RatePlan json.RawMessage `json:"rate_plan"`
	}
	// Hotel contains all the info related to the hotel
	Hotel struct {
		ID        string     `json:"hotel_id" gorm:"primaryKey;size:8"`
		Object    string     `json:"-"`
		Rooms     []Room     `json:"-" gorm:"foreignKey:HotelID"`
		RatePlans []RatePlan `json:"-" gorm:"foreignKey:HotelID"`
	}
	// Room is an entity belonging to a hotel
	Room struct {
		ID      uint   `gorm:"primaryKey"`
		HotelID string `gorm:"size:8;not null"`
		Object  string
	}
	// RatePlan ...
	RatePlan struct {
		ID      uint   `gorm:"primaryKey"`
		HotelID string `gorm:"size:8;not null"`
		Object  string
	}
)
