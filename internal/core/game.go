package core

import (
	"errors"
	"slices"
	"time"

	"github.com/google/uuid"
)

type Game struct {
  ID          string        `json:"id"`
  Format      GameFormat    `json:"format"`
  Status      GameStatus    `json:"status"`
  Host        string        `json:"host"`
  MinPlayers  int           `json:"min_players"`
  MaxPlayers  int           `json:"max_players"`
  Players     []string      `json:"players"`
  Winner      *string       `json:"winner"`
  StartedAt   *time.Time    `json:"started_at"`
  FinishedAt  *time.Time    `json:"finished_at"`
}

func NewGame(gameFormat GameFormat, minPlayers, maxPlayers int, host string) *Game {
  g := &Game{
    ID: uuid.New().String(),
    Format: gameFormat,
    Host: host,
    Status: Created,
    MinPlayers: minPlayers,
    MaxPlayers: maxPlayers,
    Players: make([]string, 0, maxPlayers),
    Winner: nil,
    StartedAt: nil,
    FinishedAt: nil,
  }

  g.Players = append(g.Players, host)

  return g
}

func (g *Game) AddPlayer(pid string) error {
  if g.Status != Created && g.Status != WaitingForPlayers {
    return errors.New("Cannot add players to a game that is not accepting more players")
  }

  if len(g.Players) >= g.MaxPlayers {
    return errors.New("Game is full")
  }

  if slices.Contains(g.Players, pid) {
    return errors.New("Player is already in game")
  }

  g.Players = append(g.Players, pid)

  return nil
}

