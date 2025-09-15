package game

import (
	"errors"
	"strings"
)

type Status int

const (
	Created Status = iota
	WaitingForPlayers
	Cancelled
	Started
	Abandoned
	Finished
)

var statusStrings = map[Status]string{
	Created:           "CREATED",
	WaitingForPlayers: "WAITING_FOR_PLAYERS",
	Cancelled:         "CANCELLED",
	Started:           "STARTED",
	Abandoned:         "ABANDONED",
	Finished:          "FINISHED",
}

var stringToStatus = map[string]Status{
	"CREATED":             Created,
	"WAITING_FOR_PLAYERS": WaitingForPlayers,
	"CANCELLED":           Cancelled,
	"STARTED":             Started,
	"ABANDONED":           Abandoned,
	"FINISHED":            Finished,
}

func NewStatusFromString(s string) (Status, error) {
	if s == "" {
		return Status(0), errors.New("Empty string cannot create a game status")
	}

	upperS := strings.ToUpper(s)

	if status, exists := stringToStatus[upperS]; exists {
		return status, nil
	}

	return Status(0), errors.New("Invalid game status string:" + s)
}

func (s Status) String() string {
	if str, exists := statusStrings[s]; exists {
		return str
	}

	return "UNKNOWN"
}

func (s Status) Equals(other Status) bool {
	return s == other
}
