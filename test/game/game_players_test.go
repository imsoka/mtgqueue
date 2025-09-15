package test

import (
	"soka/mtgqueue/internal/domain/game"
	"soka/mtgqueue/internal/domain/player"
	"testing"
	"time"
)

type fixedTime struct{}

func (fixedTime) Now() time.Time { return time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC) }

func TestGame_Join_And_Leave(t *testing.T) {
	ft := fixedTime{}
	owner := player.NewID()
	g, err := game.New(ft, game.Commander, 2, owner)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if g.Status != game.Created {
		t.Fatalf("expected Created, got %s", g.Status.String())
	}

	// First join transitions to WaitingForPlayers
	p1 := player.NewID()
	if err := g.Join(ft, p1); err != nil {
		t.Fatalf("join p1: %v", err)
	}
	if g.Status != game.WaitingForPlayers {
		t.Errorf("expected WaitingForPlayers, got %s", g.Status.String())
	}
	if !g.HasPlayer(p1) || g.PlayersCount() != 1 {
		t.Errorf("expected one player p1")
	}

	// Duplicate join denied
	if err := g.Join(ft, p1); err == nil {
		t.Errorf("expected error on duplicate join")
	}

	// Second join fills the game
	p2 := player.NewID()
	if err := g.Join(ft, p2); err != nil {
		t.Fatalf("join p2: %v", err)
	}
	if !g.IsFull() {
		t.Errorf("expected game to be full")
	}

	// Further joins denied when full
	p3 := player.NewID()
	if err := g.Join(ft, p3); err == nil {
		t.Errorf("expected error when full")
	}

	// Leave works and un-fills
	if err := g.Leave(ft, p1); err != nil {
		t.Fatalf("leave p1: %v", err)
	}
	if g.IsFull() {
		t.Errorf("expected not full after leave")
	}

	// Leave non-member denied
	px := player.NewID()
	if err := g.Leave(ft, px); err == nil {
		t.Errorf("expected error leaving non-member")
	}
}

func TestGame_Join_OnlyWhenAcceptingPlayers(t *testing.T) {
	ft := fixedTime{}
	owner := player.NewID()
	g, err := game.New(ft, game.Commander, 2, owner)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Move to Started and try to join
	if err := g.AdvanceStatus(game.WaitingForPlayers); err != nil {
		t.Fatalf("advance to waiting: %v", err)
	}
	if err := g.AdvanceStatus(game.Started); err != nil {
		t.Fatalf("advance to started: %v", err)
	}

	if err := g.Join(ft, player.NewID()); err == nil {
		t.Errorf("expected error joining when started")
	}
}

func TestGame_Leave_NotAllowedWhenStarted(t *testing.T) {
  ft := fixedTime{}
  owner := player.NewID()
  g, err := game.New(ft, game.Commander, 3, owner)
  if err != nil { t.Fatalf("unexpected error: %v", err) }

  p1 := player.NewID()
  p2 := player.NewID()
  if err := g.Join(ft, p1); err != nil { t.Fatalf("join p1: %v", err) }
  if err := g.Join(ft, p2); err != nil { t.Fatalf("join p2: %v", err) }

  if err := g.AdvanceStatus(game.Started); err != nil { t.Fatalf("advance to started: %v", err) }

  // Leave should be rejected once started
  if err := g.Leave(ft, p1); err == nil {
    t.Fatalf("expected error when leaving after start")
  }
  // Players remain unchanged
  if !g.HasPlayer(p1) || !g.HasPlayer(p2) || g.PlayersCount() != 2 {
    t.Fatalf("players should remain in game after failed leave")
  }
}

func TestGame_Leave_NotAllowedInFinalStates(t *testing.T) {
  ft := fixedTime{}
  owner := player.NewID()
  g, err := game.New(ft, game.Commander, 3, owner)
  if err != nil { t.Fatalf("unexpected error: %v", err) }

  p1 := player.NewID()
  p2 := player.NewID()
  if err := g.Join(ft, p1); err != nil { t.Fatalf("join p1: %v", err) }
  if err := g.Join(ft, p2); err != nil { t.Fatalf("join p2: %v", err) }

  // Cancelled: can only be reached before start
  gc, err := game.New(ft, game.Commander, 3, owner)
  if err != nil { t.Fatalf("unexpected error: %v", err) }
  if err := gc.Join(ft, player.NewID()); err != nil { t.Fatalf("join: %v", err) }
  if err := gc.AdvanceStatus(game.Cancelled); err != nil { t.Fatalf("advance to cancelled: %v", err) }
  if err := gc.Leave(ft, owner); err == nil { t.Fatalf("expected error leaving when cancelled") }

  // Finished: start and then add winner
  if err := g.AdvanceStatus(game.Started); err != nil { t.Fatalf("advance to started: %v", err) }
  if err := g.AddWinner(ft, p1); err != nil { t.Fatalf("add winner: %v", err) }
  if err := g.Leave(ft, p2); err == nil { t.Fatalf("expected error leaving when finished") }

  // Abandoned: create new started game, then abandon
  g2, err := game.New(ft, game.Commander, 3, owner)
  if err != nil { t.Fatalf("unexpected error: %v", err) }
  p3 := player.NewID()
  if err := g2.Join(ft, p3); err != nil { t.Fatalf("join p3: %v", err) }
  if err := g2.AdvanceStatus(game.Started); err != nil { t.Fatalf("advance to started: %v", err) }
  if err := g2.AdvanceStatus(game.Abandoned); err != nil { t.Fatalf("advance to abandoned: %v", err) }
  if err := g2.Leave(ft, p3); err == nil { t.Fatalf("expected error leaving when abandoned") }
}
