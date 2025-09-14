package test

import (
	"soka/mtgqueue/internal/domain/game"
	"testing"
	"time"
)

type MockTimestamp struct{}
func (mt *MockTimestamp) Now() time.Time {
  t := time.Date(2015, 7, 5, 12, 0, 0, 0, nil)

  return t
}


func TestGame_AdvanceStatus(t *testing.T) {
  tests := []struct {
    name              string
    status1           game.Status
    status2           game.Status
    expectedError     bool
  } {
    { "advance from CREATED to WAITING_FOR_PLAYERS", game.Created, game.WaitingForPlayers, false },
    { "advance from WAITING_FOR_PLAYERS to STARTED", game.WaitingForPlayers, game.Started, false },
    { "advance from STARTED to FINISHED", game.Started, game.Finished, false },
    { "advance from CREATED to CANCELLED", game.Created, game.Cancelled, false },
    { "advance from WAITING_FOR_PLAYERS to CANCELLED", game.WaitingForPlayers, game.Cancelled, false },
    { "advance from STARTED to ABANDONED", game.Started, game.Abandoned, false },
    { "advance from CREATED to STARTED", game.Created, game.Started, true },
    { "advance from CREATED to ABANDONED", game.Created, game.Abandoned, true },
    { "advance from CREATED to FINISHED", game.Created, game.Finished, true },
    { "advance from WAITING_FOR_PLAYERS to FINISHED", game.WaitingForPlayers, game.Finished, true },
    { "advance from WAITING_FOR_PLAYERS to ABANDONED", game.WaitingForPlayers, game.Abandoned, true },
    { "advance from WAITING_FOR_PLAYERS to CREATED", game.WaitingForPlayers, game.Created, true },
    { "advance from CANCELLED to CREATED", game.Cancelled, game.Created, true },
    { "advance from CANCELLED to WAITING_FOR_PLAYERS", game.Cancelled, game.WaitingForPlayers, true },
    { "advance from CANCELLED to STARTED", game.Cancelled, game.Started, true },
    { "advance from CANCELLED to FINISHED", game.Cancelled, game.Finished, true },
    { "advance from CANCELLED to ABANDONED", game.Cancelled, game.Abandoned, true },
    { "advance from STARTED to CREATED", game.Started, game.Created, true },
    { "advance from STARTED to WAITING_FOR_PLAYERS", game.Started, game.WaitingForPlayers, true },
    { "advance from STARTED to CANCELLED", game.Started, game.Cancelled, true },
    { "advance from ABANONED to CREATED", game.Abandoned, game.Created, true },
    { "advance from ABANDONED to WAITING_FOR_PLAYERS", game.Abandoned, game.WaitingForPlayers, true },
    { "advance from ABANDONED to CANCELLED", game.Abandoned, game.Cancelled, true },
    { "advance from ABANDONED to STARTED", game.Abandoned, game.Started, true },
    { "advance from ABANDONED to FINISHED", game.Abandoned, game.Finished, true },
    { "advance from FINISHED to CREATED", game.Finished, game.Created, true },
    { "advance from FINISHED to WAITING_FOR_PLAYERS", game.Finished, game.WaitingForPlayers, true },
    { "advance from FINISHED to CANCELLED", game.Finished, game.Cancelled, true },
    { "advance from FINISHED to STARTED", game.Finished, game.Started, true },
    { "advance from FINISHED to ABANDONED", game.Finished, game.Abandoned, true },
  }

  for _ , tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      game := &game.Game{Status: tt.status1}

      err := game.AdvanceStatus(tt.status2)

      if tt.expectedError && err == nil{
        t.Errorf("Expected error but got none")
      }

      if !tt.expectedError && err != nil {
        t.Errorf("Unexcepted error: %v", err)
      }

      if !tt.expectedError && game.Status != tt.status2 {
        t.Errorf("Expected status %s, but got %s", tt.status2.String(), game.Status.String())
      }
    })
  }
}

func TestGame_Start(t *testing.T) {
  tests := []struct {
    name              string
    status            game.Status
    expectedError     bool
  }{
    { "Start game from WAITING_FOR_PLAYERS", game.WaitingForPlayers, false },
    { "Start game from CREATED", game.Created, true },
    { "Start game from CANCELLED", game.Cancelled, true },
    { "Start game from STARTED", game.Started, true },
    { "Start game from ABANDONED", game.Abandoned, true },
    { "Start game from FINISHED", game.Finished, true },
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      g := &game.Game{
        Status: tt.status,
        StartedAt: nil,
        FinishedAt: nil,
      }

      mt := MockTimestamp{}

      err := g.Start(mt.Now())

      if tt.expectedError {
        if err == nil {
          t.Errorf("Expected error but got none")
        }
      } else {
        if err != nil {
          t.Errorf("Unexpected error for status %s: %v", g.Status.String(), err)
        }
        if !game.Status.Equals(game.Started) {
          t.Errorf("Game did not start, current game status: %s", g.Status.String())
        }

        expectedTime := mt.Now()
        if g.StartedAt == nil {
          t.Errorf("StartedAt should be %v, but got nil")
        }

        if !g.StartedAt.Equals(expectedTime) {
          t.Errorf("StartedAt should be %v, but got %v", expectedTime, g.StartedAt)
        }
      }
    })
  }
}

func TestGame_Finish(t *testing.T) {
  tests := []struct {
    name            string
    status          game.Status
    expectedError   bool
  } {
    { "Finish game from STARTED", game.Started, false },
    { "Finish game from CREATED", game.Created, true },
    { "Finish game from WAITING_FOR_PLAYERS", game.WaitingForPlayers, true },
    { "Finish game from CANCELLED", game.Cancelled, true },
    { "Finish game from ABANDONED", game.Abandoned, true },
    { "Finish game from FINISHED", game.Finished, true },
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {

      mt := MockTimestamp{}

      g := &game.Game{
        Status: tt.status,
        StartedAt: nil,
        FinishedAt: nil,
      }
      err := g.Finish(mt.Now())

      if tt.expectedError {
        if err == nil {
          t.Errorf("Expected error but got none")
        }
      } else {
        if err != nil {
          t.Errorf("Unexpected error for status %s: %v", g.Status.String(), err)
        }

        if !g.Status.Equals(g.Finished) {
          t.Errorf("Expected game to be finished, but status is %s", g.Status.String())
        }

        expectedTime := mt.Now()
        if g.FinishedAt == nil {
          t.Errorf("g.FinishedAt should be %v, but got nil", expectedTime)
        }

        if !g.FinishedAt.Equals(expectedTime) {
          t.Errorf("g.FinishedAt should be %v, but got %v", expectedTime, g.FinishedAt)
        }
      }
    })
  }
}
