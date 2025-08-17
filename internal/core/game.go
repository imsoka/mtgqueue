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
  pc := len(g.Players)

  error := g.canAddPlayers(pid, pc)

  if error != nil {
    return error
  }

  g.Players = append(g.Players, pid)

  g.updateGameStatus()


  return nil
}

func (g *Game) RemovePlayer(pid string) error {

  error := g.canRemovePlayer(pid)

  if error != nil {
    return error
  }

  playerIndex := slices.Index(g.Players, pid)
  if playerIndex == -1 {
    return errors.New("Player not found in this game")
  }

  g.Players = slices.Delete(g.Players, playerIndex, playerIndex+1)

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

  //TODO: increase players gamesPlayed count

  return nil
}

func (g *Game) Finish(winner string) error {
  now := time.Now()

  if g.Status != Active {
    return errors.New("Cannot finish a non-active game")
  }

  if g.Status == Finished {
    return errors.New("Cannot finished an already finished game")
  }

  g.Status = Finished
  g.Winner = winner
  g.FinishedAt = &now

  //TODO: Assign win to player

  return nil
}

func (g *Game) Cancel() error {
  now := time.Now()

  if g.Status != Created && g.Status != WaitingForPlayers && g.Status != Ready {
    return errors.New("Only not started games can be cancelled")
  }

  g.Status = Cancelled
  g.FinishedAt = &now

  return nil
}

func (g *Game) Abandon() error {
  now := time.Now()

  if g.Status != Active {
    return errors.New("Only active games can be abandoned")
  }

  g.Status = Abandoned
  g.FinishedAt = &now
}

func (g *Game) canRemovePlayer(pid string) error {
  nonRemovableStates := []GameStatus{Active, Finished, Cancelled, Abandoned}

  if !slices.Contains(nonRemovableStates, g.Status) {
    return errors.New("Cannot remove a player from a game that's active or finished")
  }

  return nil
}

func (g *Game) canAddPlayers(pid string, pc int) error {
  nonAppendableStates := []GameStatus{
      Finished,
      Cancelled,
      Ready,
      Active,
      Abandoned,
    }

  if pc >= g.MaxPlayers {
    return errors.New("Game is full")
  }

  if slices.Contains(g.Players, pid) {
    return errors.New("Player is already in game")
  }

  if(slices.Contains(nonAppendableStates, g.Status)) {
    return errors.New("Only can add players to game with status 'created' and 'waiting-for-players'")
  }

  return nil
}

func (g *Game) updateGameStatus() {
  pc := len(g.Players)

  switch {
	  case pc == 0:
		  g.Status = Cancelled
    case pc == 1:
		  g.Status = WaitingForPlayers
	  case pc == g.MaxPlayers:
		  g.Status = Ready
    case g.Winner != nil
      g.Status = Finished
	  default:
		  g.Status = WaitingForPlayers
	}
}
