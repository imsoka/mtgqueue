package test

import (
	"soka/mtgqueue/internal/domain/player"
	"testing"

	"github.com/google/uuid"
)

func TestPlayerID_CanCreatePlayerIDs(t *testing.T) {
	playerId := player.NewID()

	if _, err := uuid.Parse(playerId.String()); err != nil {
		t.Errorf("PlayerID gotten from constructor is invalid")
	}
}

func TestPlayerID_Constructor(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Create basic ID"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := player.NewID()
			if _, err := uuid.Parse(id.String()); err != nil {
				t.Errorf("Failed to create player ID: %v", err)
			}
		})
	}
}

func TestPlayerID_CanCreatePlayerIDsFromString(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedError bool
	}{
		{"Create from valid uuid", uuid.NewString(), false},
		{"Create from empty string", "", true},
		{"Create from invalid uuid", "INVALID", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := player.NewIDFromString(tt.input)

			if tt.expectedError {
				if err == nil {
					t.Errorf("Expected error for input %s, but got none", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for input %s: %v", tt.input, err)
				}

				if id.String() == "" {
					t.Errorf("Expected valid uuid, got empty string")
				}
			}
		})

	}

}

func TestPlayerID_EqualsWithConstructors(t *testing.T) {
	first := player.NewID()
	second := player.NewID()

	if first.Equals(second) {
		t.Errorf("Two newly created PlayerIDs should never be equals to each other")
	}
}

func TestPlayerID_EqualsWithSameUuidShouldReturnTrue(t *testing.T) {
	u := uuid.NewString()
	var id, other *player.ID
	var err error
	id, err = player.NewIDFromString(u)
	if err != nil {
		t.Errorf("Failed creating first PlayerID from given uuid %s", u)
	}
	other, err = player.NewIDFromString(u)
	if err != nil {
		t.Errorf("Failed creating second PlayerID from given uuid %s", u)
	}

	if !id.Equals(other) {
		t.Errorf("PlayerIDs created from the same UUID should be equals")
	}
}
