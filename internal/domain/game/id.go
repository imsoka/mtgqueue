package game

import (
	"encoding/json"
	"errors"
	"soka/mtgqueue/internal/domain"

	"github.com/google/uuid"
)


type ID struct {
  value     string
}

func NewID() *ID {
  return &ID{value: uuid.NewString()}
}

func NewIDFromString(s string) (*ID, error) {
  if s == "" {
    return nil, errors.New("String to create ID cannot be empty")
  }

  if _, err := uuid.Parse(s); err != nil {
    return nil, errors.New("String cannot be parsed as an uuid")
  }

  return &ID{value: s}, nil
}

func (id *ID) String() string {
  return id.value
}

func (id *ID) MarshalJSON() ([]byte, error) {
  return json.Marshal(id.value)
}

func (id *ID) UnmarshalJSON(data []byte) error {
  var s string

  if err := json.Unmarshal(data, &s); err != nil {
    return err
  }

  if _,err := uuid.Parse(s); err != nil {
    return err
  }

  id.value = s
  return nil
}

func (id *ID) Equals(other domain.Identifier) bool {

  //Transform Identifier to ID
  otherID, ok := other.(*ID)
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

func (id *ID) IsEmpty() bool {
  return id.value == ""
}

func (id *ID) IsNil() bool {
  return id == nil
}

func (id *ID) Copy() *ID {
  if id == nil {
    return nil
  }

  return &ID{value: id.value}
}

