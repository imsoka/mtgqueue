package player

import "soka/mtgqueue/internal/domain/ids"

// Re-export the shared PlayerID type in the player package for ergonomics.
type ID = ids.PlayerID

func NewID() *ID                            { return ids.NewPlayerID() }
func NewIDFromString(s string) (*ID, error) { return ids.NewPlayerIDFromString(s) }
