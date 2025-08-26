package game

import (
	"encoding/json"
	"errors"
	"soka/mtgqueue/internal/domain"

	"github.com/google/uuid"
)


type GameID struct {
  value     string
}

func NewGameID() *GameID {
  return &GameID{value: uuid.NewString()}
}

func NewGameIDFromString(s string) (*GameID, error) {
  if s == "" {
    return nil, errors.New("String to create GameID cannot be empty")
  }

  if _, err := uuid.Parse(s); err != nil {
    return nil, errors.New("String cannot be parsed as an uuid")
  }

  return &GameID{value: s}, nil
}

func (gid *GameID) String() string {
  return gid.value
}

func (gid *GameID) MarshalJSON() ([]byte, error) {
  return json.Marshal(gid.value)
}

func (gid *GameID) UnmarshalJSON(data []byte) error {
  var s string

  if err := json.Unmarshal(data, &s); err != nil {
    return err
  }

  if _,err := uuid.Parse(s); err != nil {
    return err
  }

  gid.value = s
  return nil
}

func (gid *GameID) Equals(other domain.Identifier) bool {

  //Transform Identifier to GameID
  otherGameID, ok := other.(*GameID)
  if !ok {
    return false
  }

  if gid == nil && otherGameID == nil {
    return true
  }

  if gid == nil || otherGameID == nil {
    return false
  }

  return gid.value == otherGameID.value
}

func (gid *GameID) IsEmpty() bool {
  return gid.value == ""
}

func (gid *GameID) IsNil() bool {
  return gid == nil
}

func (gid *GameID) Copy() *GameID {
  if gid == nil {
    return nil
  }

  return &GameID{value: gid.value}
}

