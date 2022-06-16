package versions

import (
	"encoding/json"
	"net/http"
)

const DataLink = "https://github.com/Ayfri/minecraft-wiki-scrapper/raw/master/out/versions.json"

type Versions []Version

func FetchVersions() (Versions, error) {
	resp, err := http.Get(DataLink)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data Versions
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func (v *Versions) GetVersion(name string) *Version {
	for _, v := range *v {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

func (v *Versions) GetLatestVersion() *Version {
	var latest Version
	for _, v := range *v {
		if v.ReleaseDate.After(latest.ReleaseDate) {
			latest = v
		}
	}
	return &latest
}

func (v *Versions) SplitByMenu() map[string]Versions {
	return SplitVersionsByMenu(v)
}
