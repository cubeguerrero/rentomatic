package use_cases

import "github.com/cubeguerrero/rentomatic/domain"

func RoomList(r roomRepoInterface) []*domain.Room {
	return r.List()
}
