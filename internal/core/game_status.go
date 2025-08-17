package core

type GameStatus int

const (
	Created GameStatus = iota
	WaitingForPlayers
	Active
	Finished
	Cancelled
	Abandoned
  Ready
)

var gameStatusNames = map[GameStatus]string {
	Created:            "created",
	WaitingForPlayers:  "waiting_for_players",
	Active:             "active",
	Finished:           "finished",
	Cancelled:          "cancelled",
	Abandoned:          "abandoned",
  Ready:              "ready",
}

func (gs GameStatus) String() string {
	if name, exists := gameStatusNames[gs]; exists {
		return name
	}
	return "unknown"
}

