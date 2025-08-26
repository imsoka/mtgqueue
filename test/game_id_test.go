package test

import (
	"soka/mtgqueue/internal/domain/game"
	"testing"

	"github.com/google/uuid"
)

func TestGameID_CanCreateGameIDs(t *testing.T) {
  gameId := game.NewGameID()

  if _, err := uuid.Parse(gameId.String()); err != nil {
    t.Errorf("GameID gotten from constructor is invalid")
  }
}

func TestGameID_CanCreateGameIDsFromString(t *testing.T) {
  uuid := uuid.NewString()

  if _, err := game.NewGameIDFromString(uuid); err != nil {
    t.Errorf("Cannot create GameID from a valid uuid")
  }
}

func TestGameID_NewGameIDFromStringShouldFailIfEmptyString(t *testing.T) {
  s := ""

  if _, err := game.NewGameIDFromString(s); err == nil {
    t.Errorf("Shouldn't be able to create a GameID from an empty string")
  }
}

func TestGameID_TestNewGameIDFromStringShouldFailIfInvalidUuidIsGiven(t *testing.T) {
  s := "superinvaliduuid"

  if _, err := game.NewGameIDFromString(s); err == nil {
    t.Errorf("Shouldn't be able to create a GameID from an invalid uuid string")
  }
}

func TestGameID_EqualsWithConstructors(t *testing.T) {
  first := game.NewGameID()
  second := game.NewGameID()

  if first.Equals(second) {
    t.Errorf("Two newly created GameIDs should never be equals to each other")
  }
}

func TestGameID_EqualsWithSameUuidShouldReturnTrue(t *testing.T) {
  uuid := uuid.NewString()
  var gid, other *game.GameID
  var err error
  gid, err = game.NewGameIDFromString(uuid)
  if err != nil {
    t.Errorf("Failed creating first GameID from given uuid %s", uuid)
  }
  other, err = game.NewGameIDFromString(uuid)
  if err != nil {
    t.Errorf("Failed creating second GameID from given uuid %s", uuid)
  }

  if !gid.Equals(other) {
    t.Errorf("GameIDs created from the same UUID should be equals")
  }
}
