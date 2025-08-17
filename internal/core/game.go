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
  if g.canAddPlayersByStatus() {
    return errors.New("Cannot add players to the game if it's active, ready or finished")
  }

  pc := len(g.Players)


  if pc >= g.MaxPlayers {
    return errors.New("Game is full")
  }

  if slices.Contains(g.Players, pid) {
    return errors.New("Player is already in game")
  }

  g.Players = append(g.Players, pid)

  pc = len(g.Players)

  g.updateGameStatus()


  return nil
}

func (g *Game) RemovePlayer(pid string) error {
  statusError := "Cannot remove player when game is active or is finished"
  if !g.canRemovePlayerByStatus() {
    return errors.New(statusError)
  }

  g.updateGameStatus()

  return nil
}

func (g *Game) Start() error {
  if g.Status != Ready && g.Status != WaitingForPlayers {
    return errors.New("Game cannot start if it isn't ready")
  }

  if len(g.Players) < g.MinPlayers {
    return errors.New("If player count is below minPlayers, the game cannot start")
  }

  now := time.Now()
  g.StartedAt = &now
  g.Status = Active

  return nil
}

func (g *Game) canRemovePlayerByStatus() bool {
  nonRemovableStates := []GameStatus{Active, Finished, Cancelled, Abandoned}

  return !slices.Contains(nonRemovableStates, g.Status)
}

func (g *Game) canAddPlayersByStatus() bool {
  nonAppendableStates := []GameStatus{
      Finished,
      Cancelled
      Ready
      Active
      Abandoned
    }

  return !slices.Contains(nonAppendableStates, g.Status)
}

func (g *Game) updateGameStatus() {
  playerCount := len(g.Players)

  switch {
	  case playerCount == 0:
		  g.Status = Cancelled
    case playerCount == 1:
		  g.Status = WaitingForPlayers
	  case playerCount == g.MaxPlayers:
		  g.Status = Ready
	  default:
		  g.Status = WaitingForPlayers
	}
}
