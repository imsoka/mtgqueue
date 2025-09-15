package game

import "soka/mtgqueue/internal/domain/ids"

// Re-export the shared GameID type in the game package for ergonomics.
type ID = ids.GameID

func NewID() *ID                            { return ids.NewGameID() }
func NewIDFromString(s string) (*ID, error) { return ids.NewGameIDFromString(s) }
