package versions

import "time"

type Version struct {
	Name        string    `json:"name"`
	ReleaseDate time.Time `json:"releaseTime,omitempty"`
	Description string    `json:"description,omitempty"`
}
