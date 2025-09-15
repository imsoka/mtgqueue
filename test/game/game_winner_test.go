package test

import (
	"soka/mtgqueue/internal/domain/game"
	"soka/mtgqueue/internal/domain/player"
	"testing"
	"time"
)

type fixedWinnerTime struct{}

func (fixedWinnerTime) Now() time.Time { return time.Date(2022, 3, 4, 5, 6, 7, 0, time.UTC) }

func TestGame_AddWinner_HappyPath_FinishesGame(t *testing.T) {
	ft := fixedWinnerTime{}
	owner := player.NewID()
	g, err := game.New(ft, game.Commander, 4, owner)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Prepare: add players and start the game
	p1 := player.NewID()
	p2 := player.NewID()
	if err := g.Join(ft, p1); err != nil {
		t.Fatalf("join p1: %v", err)
	}
	if err := g.Join(ft, p2); err != nil {
		t.Fatalf("join p2: %v", err)
	}
	if err := g.AdvanceStatus(game.Started); err != nil {
		t.Fatalf("advance to started: %v", err)
	}

	// Act: add winner
	if err := g.AddWinner(ft, p1); err != nil {
		t.Fatalf("add winner: %v", err)
	}

	if g.Status != game.Finished {
		t.Errorf("expected Finished, got %s", g.Status.String())
	}
	if g.WinnerID == nil || !g.WinnerID.Equals(p1) {
		t.Errorf("expected winner to be p1")
	}
}

func TestGame_AddWinner_OnlyWhenStarted(t *testing.T) {
	ft := fixedWinnerTime{}
	owner := player.NewID()
	g, err := game.New(ft, game.Commander, 4, owner)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	p := player.NewID()
	if err := g.Join(ft, p); err != nil {
		t.Fatalf("join p: %v", err)
	}

	// Not started yet
	if err := g.AddWinner(ft, p); err == nil {
		t.Fatalf("expected error adding winner before started")
	}
}

func TestGame_AddWinner_PlayerMustBeInGame(t *testing.T) {
	ft := fixedWinnerTime{}
	owner := player.NewID()
	g, err := game.New(ft, game.Commander, 4, owner)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Start but winner isn't a participant
	if err := g.AdvanceStatus(game.WaitingForPlayers); err != nil {
		t.Fatalf("advance: %v", err)
	}
	if err := g.AdvanceStatus(game.Started); err != nil {
		t.Fatalf("advance: %v", err)
	}

	outsider := player.NewID()
	if err := g.AddWinner(ft, outsider); err == nil {
		t.Fatalf("expected error adding non-player as winner")
	}
}

func TestGame_AddWinner_CanOverwrite(t *testing.T) {
	ft := fixedWinnerTime{}
	owner := player.NewID()
	g, err := game.New(ft, game.Commander, 4, owner)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	p1 := player.NewID()
	p2 := player.NewID()
	if err := g.Join(ft, p1); err != nil {
		t.Fatalf("join p1: %v", err)
	}
	if err := g.Join(ft, p2); err != nil {
		t.Fatalf("join p2: %v", err)
	}
	if err := g.AdvanceStatus(game.Started); err != nil {
		t.Fatalf("advance to started: %v", err)
	}

    if err := g.AddWinner(ft, p1); err != nil {
        t.Fatalf("add winner: %v", err)
    }
    if err := g.AddWinner(ft, p2); err != nil {
        t.Fatalf("overwriting winner should be allowed: %v", err)
    }
    if g.WinnerID == nil || !g.WinnerID.Equals(p2) {
        t.Fatalf("expected winner to be overwritten to p2")
    }
    if g.Status != game.Finished {
        t.Fatalf("game should remain finished after overwrite")
    }
}
