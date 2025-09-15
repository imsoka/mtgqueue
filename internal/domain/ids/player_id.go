package ids

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"soka/mtgqueue/internal/domain"
)

type PlayerID struct {
	value string
}

func NewPlayerID() *PlayerID {
	return &PlayerID{value: uuid.NewString()}
}

func NewPlayerIDFromString(s string) (*PlayerID, error) {
	if s == "" {
		return nil, errors.New("String to create ID cannot be empty")
	}
	if _, err := uuid.Parse(s); err != nil {
		return nil, errors.New("String cannot be parsed as an uuid")
	}
	return &PlayerID{value: s}, nil
}

func (id *PlayerID) String() string {
	return id.value
}

func (id *PlayerID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.value)
}

func (id *PlayerID) UnmarshalJSON(data []byte) error {
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

func (id *PlayerID) Equals(other domain.Identifier) bool {
	otherID, ok := other.(*PlayerID)
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

func (id *PlayerID) IsEmpty() bool { return id.value == "" }
func (id *PlayerID) IsNil() bool   { return id == nil }

func (id *PlayerID) Copy() *PlayerID {
	if id == nil {
		return nil
	}
	return &PlayerID{value: id.value}
}
