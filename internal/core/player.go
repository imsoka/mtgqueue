package core

import ()

type Player struct {
  ID            string    `json:"id"`
  Username      string    `json:"username"`
  GamesPlayed   int       `json:"games_played"`
  GamesTied     int       `json:"games_tied"`
  GamesWon      int       `json:"games_won"`
  Rating        *int      `json:"rating"`
}


func NewPlayer(id, username string) *Player {
  return &Player{
    ID: id,
    Username: username,
    GamesPlayed: 0,
    GamesTied: 0,
    GamesWon: 0,
    Rating: nil,
  }
}
