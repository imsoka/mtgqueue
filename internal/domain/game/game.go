package game

import (
	"errors"
	"soka/mtgqueue/internal/domain"
	"soka/mtgqueue/internal/domain/ids"
	"time"
)

// Game represents a game lifecycle and the players in it.
type Game struct {
	ID         *ID
	Format     Format
	Status     Status
	MaxPlayers int
	OwnerID    *ids.PlayerID
	PlayerIDs  []*ids.PlayerID
	WinnerID   *ids.PlayerID
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// New creates a new game with the given format and capacity.
func New(t domain.Timestamp, f Format, mp int, o *ids.PlayerID) (*Game, error) {
	if !f.IsValid() {
		return nil, errors.New("invalid format")
	}
	if mp <= 0 {
		return nil, errors.New("max players must be > 0")
	}
	if o == nil || o.IsEmpty() {
		return nil, errors.New("owner id required")
	}
	now := t.Now()
	return &Game{
		ID:         NewID(),
		Format:     f,
		Status:     Created,
		MaxPlayers: mp,
		OwnerID:    o,
		PlayerIDs:  []*ids.PlayerID{},
		CreatedAt:  now,
		UpdatedAt:  now,
	}, nil
}

// AdvanceStatus attempts to transition the game to the given status.
// It enforces valid transitions and returns an error otherwise.
func (g *Game) AdvanceStatus(next Status) error {
	allowed := false
	switch g.Status {
	case Created:
		allowed = (next == WaitingForPlayers || next == Cancelled)
	case WaitingForPlayers:
		allowed = (next == Started || next == Cancelled)
	case Started:
		allowed = (next == Finished || next == Abandoned)
	case Abandoned, Finished, Cancelled:
		allowed = false
	default:
		allowed = false
	}

	if !allowed {
		return errors.New("invalid status transition")
	}

	g.Status = next
	return nil
}

// PlayersCount returns the number of players currently in the game.
func (g *Game) PlayersCount() int { return len(g.PlayerIDs) }

// IsFull reports whether the game reached its capacity.
func (g *Game) IsFull() bool { return len(g.PlayerIDs) >= g.MaxPlayers }

// HasPlayer checks if a player is already in the game by ID string.
func (g *Game) HasPlayer(pid *ids.PlayerID) bool {
	for _, existing := range g.PlayerIDs {
		if existing.Equals(pid) {
			return true
		}
	}
	return false
}

// Join attempts to add a player to the game.
// Allowed only in Created or WaitingForPlayers and when not full.
func (g *Game) Join(t domain.Timestamp, pid *ids.PlayerID) error {
	if pid == nil || pid.IsEmpty() {
		return errors.New("player id required")
	}
	if !(g.Status == Created || g.Status == WaitingForPlayers) {
		return errors.New("game not accepting players")
	}
	if g.HasPlayer(pid) {
		return errors.New("player already joined")
	}
	if g.IsFull() {
		return errors.New("game is full")
	}
	g.PlayerIDs = append(g.PlayerIDs, pid.Copy())
	if g.Status == Created {
		g.Status = WaitingForPlayers
	}
	g.UpdatedAt = t.Now()
	return nil
}

// Leave removes a player from the game.
// If the game is waiting for players and becomes empty, it returns to Created.
func (g *Game) Leave(t domain.Timestamp, pid *ids.PlayerID) error {
    if pid == nil || pid.IsEmpty() {
        return errors.New("player id required")
    }
    if !(g.Status == Created || g.Status == WaitingForPlayers) {
        return errors.New("players can only leave before the game starts")
    }
	idx := -1
	for i, existing := range g.PlayerIDs {
		if existing.Equals(pid) {
			idx = i
			break
		}
	}
	if idx == -1 {
		return errors.New("player not in game")
	}
	g.PlayerIDs = append(g.PlayerIDs[:idx], g.PlayerIDs[idx+1:]...)
	if g.Status == WaitingForPlayers && len(g.PlayerIDs) == 0 {
		g.Status = Created
	}
	g.UpdatedAt = t.Now()
	return nil
}

// AddWinner records the winner of the game and finishes it.
// Allowed only when the game is Started and the winner is a participant.
// A winner can also be Overwritten
func (g *Game) AddWinner(t domain.Timestamp, pid *ids.PlayerID) error {
	if pid == nil || pid.IsEmpty() {
		return errors.New("player id required")
	}
	if !(g.Status == Started || g.Status == Finished) {
		return errors.New("winner can only be set when game is started or finished")
	}
	if !g.HasPlayer(pid) {
		return errors.New("winner must be a player in the game")
	}
	g.WinnerID = pid.Copy()
	g.Status = Finished
	g.UpdatedAt = t.Now()
	return nil
}
