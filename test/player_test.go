package test

import (
  "testing"
  "soka/mtgqueue/internal/core"
)

func TestNewPlayer(t *testing.T) {
  got := core.NewPlayer("1234", "Soka")

  if got.ID != "1234" {
    t.Errorf("got ID %s and wanted %s", got.ID, "1234")
  }

  if got.Username != "Soka" {
    t.Errorf("got username %s and wanted %s", got.Username, "Soka")
  }

  if got.GamesPlayed != 0 {
    t.Errorf("got GamesPlayed %d and wanted %d", got.GamesPlayed, 0)
  }

  if got.GamesWon != 0 {
    t.Errorf("got GamesWon %d and wanted %d", got.GamesWon, 0)
  }

  if got.GamesTied != 0 {
    t.Errorf("got GamesTied %d and wanted %d", got.GamesTied, 0)
  }

  if got.Rating != nil {
    t.Errorf("default Rating should be nil")
  }
}
