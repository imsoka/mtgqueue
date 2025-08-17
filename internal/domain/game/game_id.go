package game

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type GameID struct {
  value string
}

func NewGameID() GameID {
  return GameID{value: uuid.New().String()}
}

func (i GameID) GameIDFromString(s string) (GameID, error) {
  if s == "" {
    return GameID{}, errors.New("gameID cannot be empty")
  }

  if _, err := uuid.Parse(s); err != nil {
    return GameID{}, fmt.Errorf("Invalid UUID format %w", err)
  }

  return GameID{value: s}, nil
}

func (i GameID) String() string {
  return i.value
}

func (i GameID) Equals(other GameID) bool {
   return i.value == other.value
}

func (i GameID) IsEmpty() bool {
  return i.value == ""
}

func (i GameID) MarshalJSON() ([]byte, error) {
  return json.Marshal(i.value)
}
