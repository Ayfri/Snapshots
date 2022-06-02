package versions

import (
	"encoding/json"
	"time"
)

type Version struct {
	Name        string    `json:"name"`
	ReleaseDate time.Time `json:"releaseTime,omitempty"`
	Description string    `json:"description,omitempty"`
}

func (v *Version) UnmarshalJSON(data []byte) error {
	var version struct {
		Name        string `json:"name"`
		ReleaseDate int64  `json:"releaseTime,omitempty"`
		Description string `json:"description,omitempty"`
	}
	if err := json.Unmarshal(data, &version); err != nil {
		return err
	}

	v.Name = version.Name
	v.ReleaseDate = time.Unix(version.ReleaseDate, 0)
	v.Description = version.Description
	return nil
}

func (v Version) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name        string `json:"name"`
		ReleaseDate int64  `json:"releaseTime"`
		Description string `json:"description"`
	}{
		Name:        v.Name,
		ReleaseDate: v.ReleaseDate.Unix(),
		Description: v.Description,
	})
}
