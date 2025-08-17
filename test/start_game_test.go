package test

import (
	"soka/mtgqueue/internal/core"
	"testing"
)


func TestTryingToStartInvalidGameShouldGiveError (t *testing.T) {
  g := setup()

  error := g.Start()

  if error != nil {
    t.Errorf("Game shouldn't be able to start if it isn't ready")
  }
}

func TestTryingToStartWithoutMinimunPlayersShouldGiveError (t *testing.T) {
  g := setup()

  error := g.Start()


  if error != nil {
    t.Errorf("Game shouldn't be able to start if player count is below minPlayers")
  }
}


