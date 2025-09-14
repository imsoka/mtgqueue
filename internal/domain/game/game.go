package game

import (
	"errors"
	"soka/mtgqueue/internal/domain"
)

type Game struct {
  Status          Status
  StartedAt       *time.Time
  FinishedAt      *time.Time
}

func NewGame(host &player.Player) *Game {
  return &Game{
    Status:       Created,
    StartedAt:    nil,
    FinishedAt:   nil,
  }
}

func (g *Game) Start() error {
  err := g.start(nil)

  return err
}

func (g *Game) StartWithTimestamp(ts domain.Timestamp) error {
  err := g.start(ts)

  return err
}

func (g *Game) start(ts *domain.Timestamp) error {
  return errors.New("Not yet implemented")
}
