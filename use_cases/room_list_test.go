package use_cases_test

import (
	"reflect"
	"testing"

	"github.com/google/uuid"

	"github.com/cubeguerrero/rentomatic/domain"
	"github.com/cubeguerrero/rentomatic/use_cases"
)

type mockRepo struct{}

func (mr *mockRepo) List() []*domain.Room {
	room1 := domain.NewRoom(uuid.New(), 215, 39, -0.09998975, 51.75436293)
	room2 := domain.NewRoom(uuid.New(), 405, 66, 0.18228006, 51.74640997)
	room3 := domain.NewRoom(uuid.New(), 56, 60, 0.27891577, 51.45994069)
	room4 := domain.NewRoom(uuid.New(), 93, 48, 0.33894476, 51.39916678)
	return []*domain.Room{room1, room2, room3, room4}
}

func TestRoomListWithoutParameters(t *testing.T) {
	repo := &mockRepo{}

	result := use_cases.RoomList(repo)

	gotCodes := []uuid.UUID{}
	for _, r := range result {
		gotCodes = append(gotCodes, r.Code)
	}

	expectedCodes := []uuid.UUID{}
	for _, r := range repo.List() {
		expectedCodes = append(expectedCodes, r.Code)
	}

	if reflect.DeepEqual(gotCodes, expectedCodes) {
		t.Error("Error")
	}
}
