package test

import (
	"soka/mtgqueue/internal/core"
	"testing"
)

func setup() *core.Game {
  g := core.NewGame(core.Commander, 4, 4, "1234")

  return g
}

func TestAddingPlayerToFullGameShouldGiveError(t *testing.T) {
  g := setup()

  g.Players = append(g.Players, "4321")
  g.Players = append(g.Players, "3412")
  g.Players = append(g.Players, "2314")

  var error = g.AddPlayer("4231")

  if error == nil {
    t.Errorf("Game MaxPlayers is %d and %d players are inside.", g.MaxPlayers, len(g.Players))
  }
}

func TestAddingExistingPlayerShouldGiveError(t *testing.T) {
  g := setup()

  error := g.AddPlayer(g.Host)

  if error == nil {
    t.Errorf("This should have given an error. Player \"%s\" is already in game.", g.Host)
  }
}

func TestAddingPlayerToNotAcceptingGameShouldGiveError(t *testing.T) {
  g := setup()

  g.Status = core.Abandoned

  error := g.AddPlayer("3241")

  if error == nil {
    t.Errorf("This game wasn't accepting new players. Only 'created' and 'waiting-for-players' games are.")
  }
}
