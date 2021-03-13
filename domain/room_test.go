package domain_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/uuid"

	"github.com/cubeguerrero/rentomatic/domain"
)

func TestNewRoom(t *testing.T) {
	code := uuid.New()
	room := domain.NewRoom(
		code, 200, 10, 0.09998975, 51.75436293,
	)

	if room.Code != code {
		t.Errorf("Expected code: %s, got: %s", code, room.Code)
	}

	if room.Size != 200 {
		t.Errorf("Expected size: %d, got: %d", 200, room.Size)
	}

	if room.Price != 10 {
		t.Errorf("Expected price: %d, got: %d", 10, room.Price)
	}

	if room.Longitude != 0.09998975 {
		t.Errorf("Expected longitude: %f, got: %f", 0.09998975, room.Longitude)
	}

	if room.Latitude != 51.75436293 {
		t.Errorf("Expected latitude: %f, got: %f", 51.75436293, room.Latitude)
	}
}

func TestRoomSerialization(t *testing.T) {
	code := uuid.New()
	room := domain.NewRoom(
		code, 200, 10, 0.09998975, 51.75436293,
	)
	expectedJSONStringFmt := `{"code":"%s","size":200,"price":10,"longitude":0.09998975,"latitude":51.75436293}`
	expectedJSONString := fmt.Sprintf(expectedJSONStringFmt, code)
	got, _ := json.Marshal(room)

	if expectedJSONString != string(got) {
		t.Errorf("Expected: %v, Got: %v", expectedJSONString, string(got))
	}
}
