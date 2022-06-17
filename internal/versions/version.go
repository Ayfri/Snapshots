package versions

import (
	"VersionCraft2/internal/snapshots"
	"encoding/json"
	"regexp"
	"strings"
	"time"
)

type Version struct {
	Name        string               `json:"name"`
	ReleaseDate time.Time            `json:"releaseTime,omitempty"`
	Description string               `json:"description,omitempty"`
	Snapshots   []snapshots.Snapshot `json:"snapshots,omitempty"`
}

func (v *Version) UnmarshalJSON(data []byte) error {
	var version struct {
		Name        string               `json:"name"`
		ReleaseDate int64                `json:"releaseTime,omitempty"`
		Description string               `json:"description,omitempty"`
		Snapshots   []snapshots.Snapshot `json:"snapshots,omitempty"`
	}
	if err := json.Unmarshal(data, &version); err != nil {
		return err
	}

	v.Name = version.Name
	v.ReleaseDate = time.Unix(version.ReleaseDate, 0)
	v.Description = version.Description
	v.Snapshots = version.Snapshots

	return nil
}

func (v Version) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name        string               `json:"name"`
		ReleaseDate int64                `json:"releaseTime"`
		Description string               `json:"description"`
		Snapshots   []snapshots.Snapshot `json:"snapshots"`
	}{
		Name:        v.Name,
		ReleaseDate: v.ReleaseDate.Unix(),
		Description: v.Description,
		Snapshots:   v.Snapshots,
	})
}

func (v *Version) SiteName() string {
	name := v.Name
	name = regexp.MustCompile("Java Edition (?P<link>.+?)").ReplaceAllString(name, "$1")
	name = strings.ReplaceAll(name, " updates", "")

	name = regexp.MustCompile("[\\s.]").ReplaceAllString(name, "-")
	name = strings.ToLower(name)
	return name
}

func (v *Version) SiteURL() string {
	return "/versions/" + v.SiteName()
}

func (v *Version) GetSnapshotsByType(t snapshots.SnapshotType) []snapshots.Snapshot {
	var snapshots []snapshots.Snapshot
	for _, snapshot := range v.Snapshots {
		if snapshot.Type == t {
			snapshots = append(snapshots, snapshot)
		}
	}
	return snapshots
}

func (v *Version) GetSnapshotsCountByType(t snapshots.SnapshotType) int {
	var count int
	for _, snapshot := range v.Snapshots {
		if snapshot.Type == t {
			count++
		}
	}

	return count
}

func (v *Version) GetSnapshotsByTypes() (snapshots []snapshots.Snapshot) {
	for _, snapshot := range v.Snapshots {
		snapshots = append(snapshots, snapshot)
	}

	return
}
