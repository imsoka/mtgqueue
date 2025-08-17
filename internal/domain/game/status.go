package game

type Status int

const (
	Created Status = iota
	WaitingForPlayers
	Active
	Finished
	Cancelled
	Abandoned
  Ready
)

var gameStatusNames = map[Status]string {
	Created:            "created",
	WaitingForPlayers:  "waiting_for_players",
	Active:             "active",
	Finished:           "finished",
	Cancelled:          "cancelled",
	Abandoned:          "abandoned",
  Ready:              "ready",
}

func (gs Status) String() string {
	if name, exists := gameStatusNames[gs]; exists {
		return name
	}
	return "unknown"
}

