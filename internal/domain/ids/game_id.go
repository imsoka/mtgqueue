package ids

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"soka/mtgqueue/internal/domain"
)

type GameID struct {
	value string
}

func NewGameID() *GameID {
	return &GameID{value: uuid.NewString()}
}

func NewGameIDFromString(s string) (*GameID, error) {
	if s == "" {
		return nil, errors.New("String to create ID cannot be empty")
	}
	if _, err := uuid.Parse(s); err != nil {
		return nil, errors.New("String cannot be parsed as an uuid")
	}
	return &GameID{value: s}, nil
}

func (id *GameID) String() string { return id.value }

func (id *GameID) MarshalJSON() ([]byte, error) { return json.Marshal(id.value) }

func (id *GameID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if _, err := uuid.Parse(s); err != nil {
		return err
	}
	id.value = s
	return nil
}

func (id *GameID) Equals(other domain.Identifier) bool {
	otherID, ok := other.(*GameID)
	if !ok {
		return false
	}
	if id == nil && otherID == nil {
		return true
	}
	if id == nil || otherID == nil {
		return false
	}
	return id.value == otherID.value
}

func (id *GameID) IsEmpty() bool { return id.value == "" }
func (id *GameID) IsNil() bool   { return id == nil }

func (id *GameID) Copy() *GameID {
	if id == nil {
		return nil
	}
	return &GameID{value: id.value}
}
