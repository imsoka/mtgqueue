package domain

import "time"

type Timestamp interface {
	Now() time.Time
}
