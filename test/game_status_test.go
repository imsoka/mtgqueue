package test

import (
	"soka/mtgqueue/internal/domain/game"
	"testing"
)

func TestGameSatus_Constants(t *testing.T) {
  tests := []struct {
    name          string
    status        game.Status
    expected      string
  }{
    {"Created status", game.Created, "CREATED"},
    {"WaitingForPlayers status", game.WaitingForPlayers, "WAITING_FOR_PLAYERS"},
    {"Cancelled status", game.Cancelled, "CANCELLED"},
    {"Started status", game.Started, "STARTED"},
    {"Abandoned status", game.Abandoned, "ABANDONED"},
    {"Finished status", game.Finished, "FINISHED"},
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      if tt.status.String() != tt.expected {
        t.Errorf("Expected %s got %s", tt.expected, tt.status.String())
      }
    })
  }
}

func TestGameStatus_String(t *testing.T) {
  tests := []struct {
    name          string
    status        game.Status
    expected      string
  }{
    {"Created to string", game.Created, "CREATED"},
    {"WaitingForPlayers to string", game.WaitingForPlayers, "WAITING_FOR_PLAYERS"},
    {"Cancelled to string", game.Cancelled, "CANCELLED"},
    {"Started to string", game.Started, "STARTED"},
    {"Abandoned to string", game.Abandoned, "ABANDONED"},
    {"Finished to string", game.Finished, "FINISHED"},
  }

  for _,tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      result := tt.status.String()
      if result != tt.expected {
        t.Errorf("Expected %s got %s", tt.expected, result)
      }
    })
  }
}

