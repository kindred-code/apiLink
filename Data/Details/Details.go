package details

import (
	"time"
)

type Details struct {
	ProfileID      int64     `json:"profile_id"`
	Gender         string    `json:"gender"`
	Location       string    `json:"location"`
	Bio            string    `json:"bio"`
	Interests      []string  `json:"interests"`
	PhotosURLs     []string  `json:"photos_urls"`
	LastActiveTime time.Time `json:"last_active_time"`
}
