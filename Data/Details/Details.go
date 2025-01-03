package details

import (
	"time"
)

type Details struct {
	ProfileId      int64     `json:"profile_id"`
	Gender         string    `json:"gender"`
	Location       string    `json:"location"`
	Bio            string    `json:"bio"`
	LastActiveTime time.Time `json:"last_active_time"`
}
