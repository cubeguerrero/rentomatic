package repository

import (
	"github.com/cubeguerrero/rentomatic/domain"
	"github.com/google/uuid"
)

type memRepo struct {
	rooms []*domain.Room
}

func (mr *memRepo) List() []*domain.Room {
	return mr.rooms
}

func NewMemRepo(roomDicts []map[string]interface{}) (*memRepo, error) {
	rooms := make([]*domain.Room, len(roomDicts))
	for i, r := range roomDicts {
		code, err := uuid.Parse(r["code"].(string))
		if err != nil {
			return nil, err
		}
		rooms[i] = domain.NewRoom(code, r["size"].(int), r["price"].(int), r["longitude"].(float64), r["latitude"].(float64))
	}

	return &memRepo{
		rooms: rooms,
	}, nil
}
