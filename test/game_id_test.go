package test

import (
	"soka/mtgqueue/internal/domain/game"
	"testing"

	"github.com/google/uuid"
)

func TestGameID_CanCreateGameIDs(t *testing.T) {
  gameId := game.NewID()

  if _, err := uuid.Parse(gameId.String()); err != nil {
    t.Errorf("GameID gotten from constructor is invalid")
  }
}

func TestGameID_Constructor(t *testing.T) {
  tests := []struct {
    name              string
  }{
    { "Create basic ID" },
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      id := game.NewID()
      if _, err := uuid.Parse(id.String()); err != nil {
        t.Errorf("Failed to create game ID: %v", err)
      }
    })
  }
}

func TestGameID_CanCreateGameIDsFromString(t *testing.T) {
  tests := []struct {
    name            string
    input           string
    expectedError   bool
  }{
    { "Create from valid uuid", uuid.NewString(), false},
    { "Create from empty string", "", true},
    { "Create from invalid uuid", "INVALID", true},
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      id, err := game.NewIDFromString(tt.input)

      if tt.expectedError {
        if err == nil {
          t.Errorf("Expected error for input %s, but got none", tt.input)
        }
      } else {
        if err != nil {
          t.Errorf("Unexpected error for input %s: %v", tt.input, err)
        }

        if id.String() == "" {
          t.Errorf("Expected valid uuid, go empty string")
        }
      }
    })

  }

}

func TestGameID_EqualsWithConstructors(t *testing.T) {
  first := game.NewID()
  second := game.NewID()

  if first.Equals(second) {
    t.Errorf("Two newly created GameIDs should never be equals to each other")
  }
}

func TestGameID_EqualsWithSameUuidShouldReturnTrue(t *testing.T) {
  uuid := uuid.NewString()
  var id, other *game.ID
  var err error
  id, err = game.NewIDFromString(uuid)
  if err != nil {
    t.Errorf("Failed creating first GameID from given uuid %s", uuid)
  }
  other, err = game.NewIDFromString(uuid)
  if err != nil {
    t.Errorf("Failed creating second GameID from given uuid %s", uuid)
  }

  if !id.Equals(other) {
    t.Errorf("GameIDs created from the same UUID should be equals")
  }
}
