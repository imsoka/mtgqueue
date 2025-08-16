package core

import "time"

type Game struct {
  ID          string        `json:"id"`
  Players     []string      `json:"players"`
  Winner      string        `json:"winner"`
  StartedAt   time.Time     `json:"started_at"`
  FinishedAt  time.Time     `json:"finished_at"`
  MinPlayers  int           `json:"min_players"`
  MaxPlayers  int           `json:"max_players"`
