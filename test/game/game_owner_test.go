package test

import (
	"soka/mtgqueue/internal/domain/game"
	"soka/mtgqueue/internal/domain/player"
	"testing"
	"time"
)

type fixedOwnerTime struct{}

func (fixedOwnerTime) Now() time.Time { return time.Date(2021, 2, 3, 4, 5, 6, 0, time.UTC) }

func TestGame_NewWithOwner_SetsOwner(t *testing.T) {
	ft := fixedOwnerTime{}
	owner := player.NewID()
	g, err := game.New(ft, game.Commander, 4, owner)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if g.OwnerID == nil {
		t.Fatalf("expected owner to be set")
	}
	if !g.OwnerID.Equals(owner) {
		t.Errorf("owner id mismatch")
	}
}

func TestGame_New_WithoutOwner_Fails(t *testing.T) {
	ft := fixedOwnerTime{}
	g, err := game.New(ft, game.Commander, 4, nil)
	if err == nil {
		t.Fatalf("expected error when owner is nil, got game: %+v", g)
	}
}
