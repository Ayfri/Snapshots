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
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *Data) SortByDate() {
	*d = append(*d, versions.Version{})
	for i := len(*d) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if (*d)[j].ReleaseDate.Before((*d)[j+1].ReleaseDate) {
				(*d)[j], (*d)[j+1] = (*d)[j+1], (*d)[j]
			}
		}
	}
	*d = (*d)[:len(*d)-1]
}
