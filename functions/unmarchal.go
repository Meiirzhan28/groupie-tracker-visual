package functions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetAllArtist() ([]Artist, error) {
	text, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return []Artist{}, err
	}
	var art []Artist
	defer text.Body.Close()
	content, err := ioutil.ReadAll(text.Body)
	if err != nil {
		return []Artist{}, err
	}

	if err := json.Unmarshal(content, &art); err != nil {
		return []Artist{}, err
	}
	return art, nil
}

func GetOneArtist(id string) (Artist, error) {
	text, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
	if err != nil {
		return Artist{}, err
	}
	var art Artist
	defer text.Body.Close()
	content, err := ioutil.ReadAll(text.Body)
	if err != nil {
		return Artist{}, err
	}

	if err := json.Unmarshal(content, &art); err != nil {
		return Artist{}, err
	}
	return art, nil
}

func GetArtistLocations(id string) (Location, error) {
	text, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)
	if err != nil {
		return Location{}, err
	}
	defer text.Body.Close()
	content, err := ioutil.ReadAll(text.Body)
	if err != nil {
		return Location{}, err
	}
	var Loacasya Location
	if err := json.Unmarshal(content, &Loacasya); err != nil {
		return Location{}, err
	}
	return Loacasya, nil
}
