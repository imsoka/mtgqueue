package game



type Status int


const (
  Created Status = iota
  WaitingForPlayers
  Cancelled
  Abandoned
  Finished
)


var statusStrings = map[Status]string {
  Created:              "CREATED",
  WaitingForPlayers:    "WAITING_FOR_PLAYERS"
  Cancelled:            "CANCELLED"
  Started:              "STARTED"
  Abandoned:            "ABANDONED"
  Finished:             "FINISHED"
}

var stringToStatus = map[string]Status {
  "CREATED":              Created
  "WAITING_FOR_PLAYERS":  WaitingForPlayers
  "CANCELLED":            Cancelled
  "STARTED":              Started
  "ABANDONED":            Abandoned
  "FINISHED":             Finished
}
