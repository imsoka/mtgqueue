package core

type GameHistory struct {
  PlayerID        string        `json:"player_id"`
  Games           []string      `json:"games"`
}
