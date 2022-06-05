package api

import (
	"VersionCraft2/internal/versions"
	"encoding/json"
	"net/http"
)

const DataLink = "https://github.com/Ayfri/minecraft-wiki-scrapper/raw/master/out/versions.json"

type Data []versions.Version

func FetchData() (Data, error) {
	resp, err := http.Get(DataLink)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data Data
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func (d *Data) GetVersion(name string) *versions.Version {
	for _, v := range *d {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

func (d *Data) GetLatestVersion() *versions.Version {
	var latest versions.Version
	for _, v := range *d {
		if v.ReleaseDate.After(latest.ReleaseDate) {
			latest = v
		}
	}
	return &latest
}
