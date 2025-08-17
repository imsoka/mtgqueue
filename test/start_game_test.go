package test

import (
	"testing"
)


func TestTryingToStartInvalidGameShouldGiveError (t *testing.T) {
  g := setup()

  error := g.Start()

  if error != nil {
    t.Errorf("Game shouldn't be able to start if it isn't ready")
  }
}

