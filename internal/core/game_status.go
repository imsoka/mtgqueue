package core

type GameStatus int

const (
	Created GameStatus = iota
	WaitingForPlayers
	Active
	Paused
	Completed
	Cancelled
	Abandoned
)

var gameStatusNames = map[GameStatus]string{
	Created:           "created",
	WaitingForPlayers: "waiting_for_players",
	Active:            "active",
	Paused:            "paused",
	Completed:         "completed",
	Cancelled:         "cancelled",
	Abandoned:         "abandoned",
}

func (gs GameStatus) String() string {
	if name, exists := gameStatusNames[gs]; exists {
		return name
	}
	return "unknown"
}

