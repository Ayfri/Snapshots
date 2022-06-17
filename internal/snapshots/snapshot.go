package snapshots

import (
	"encoding/json"
	"regexp"
	"strings"
	"time"
)

type SnapshotType string

const (
	SnapshotTypeExperimental SnapshotType = "experimental"
	SnapshotTypeSnapshot     SnapshotType = "snapshot"
	SnapshotTypePreRelease   SnapshotType = "pre-release"
	SnapshotTypeCandidate    SnapshotType = "candidate"
	SnapshotTypeRelease      SnapshotType = "release"
)

type Snapshot struct {
	Name         string       `json:"name"`
	ReleaseDate  time.Time    `json:"releaseTime,omitempty"`
	Description  string       `json:"description"`
	DownloadJar  string       `json:"downloadClient,omitempty"`
	DownloadJSON string       `json:"downloadJSON,omitempty"`
	Type         SnapshotType `json:"snapshotType,omitempty"`
}

func (s Snapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name         string       `json:"name"`
		ReleaseDate  int64        `json:"releaseTime,omitempty"`
		Description  string       `json:"description"`
		DownloadJar  string       `json:"downloadClient,omitempty"`
		DownloadJSON string       `json:"downloadJSON,omitempty"`
		Type         SnapshotType `json:"snapshotType,omitempty"`
	}{
		Name:         s.Name,
		ReleaseDate:  s.ReleaseDate.Unix(),
		Description:  s.Description,
		DownloadJar:  s.DownloadJar,
		DownloadJSON: s.DownloadJSON,
		Type:         s.Type,
	})
}

func (s *Snapshot) UnmarshalJSON(data []byte) error {
	var snapshot struct {
		Name         string       `json:"name"`
		ReleaseDate  int64        `json:"releaseTime,omitempty"`
		Description  string       `json:"description"`
		DownloadJar  string       `json:"downloadClient,omitempty"`
		DownloadJSON string       `json:"downloadJSON,omitempty"`
		Type         SnapshotType `json:"snapshotType,omitempty"`
	}
	if err := json.Unmarshal(data, &snapshot); err != nil {
		return err
	}

	s.Name = snapshot.Name
	s.ReleaseDate = time.Unix(snapshot.ReleaseDate, 0)
	s.Description = snapshot.Description
	s.DownloadJar = snapshot.DownloadJar
	s.DownloadJSON = snapshot.DownloadJSON
	s.Type = snapshot.Type

	return nil
}
func (s *Snapshot) SiteName() string {
	name := s.Name
	name = regexp.MustCompile("Java Edition (?P<link>.+?)").ReplaceAllString(name, "$1")
	name = strings.ReplaceAll(name, " updates", "")

	name = regexp.MustCompile("[\\s.]").ReplaceAllString(name, "-")
	name = strings.ToLower(name)
	return name
}

func (s *Snapshot) SiteURL() string {
	return "/snapshots/" + s.SiteName()
}
