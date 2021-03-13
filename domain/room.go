package domain

import (
	"github.com/google/uuid"
)

type Room struct {
	Code      uuid.UUID `json:"code"`
	Size      int       `json:"size"`
	Price     int       `json:"price"`
	Longitude float64   `json:"longitude"`
	Latitude  float64   `json:"latitude"`
}

func NewRoom(code uuid.UUID, size, price int, longitude, latitude float64) *Room {
	return &Room{
		Code:      code,
		Size:      size,
		Price:     price,
		Longitude: longitude,
		Latitude:  latitude,
	}
}
