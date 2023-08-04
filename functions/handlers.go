package functions

import (
	"net/http"
	"strconv"
	"text/template"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstalbum"`
	Relations    string   `json:"relations"`
}

type Location struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Result struct {
	Art Artist
	Rel Location
}

func Home(w http.ResponseWriter, r *http.Request) {
	allArtists, err := GetAllArtist()
	if err != nil {
		Errorhandler(w, http.StatusInternalServerError)
		return
	}
	if r.URL.Path != "/" {
		Errorhandler(w, 404)
		return
	}
	if r.Method != http.MethodGet {
		Errorhandler(w, 405)
		return
	}
	text, err := template.ParseFiles("templates/index.html")
	if err != nil {
		Errorhandler(w, 500)
		return
	}
	err = text.Execute(w, allArtists)
	if err != nil {
		Errorhandler(w, 500)
		return
	}
}

func Posthandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/tracker/" {
		Errorhandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		Errorhandler(w, http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	_, err := strconv.Atoi(id)
	if err != nil {
		Errorhandler(w, http.StatusNotFound)
		return
	}
	// if idNum < 1 || idNum > 52 {
	// 	Errorhandler(w, 404)
	// 	return
	// }
	Locasya, err := GetArtistLocations(id)
	if err != nil {
		Errorhandler(w, http.StatusInternalServerError)
		return
	}
	oneArtist, err := GetOneArtist(id)
	if err != nil {
		Errorhandler(w, http.StatusInternalServerError)
		return
	}
	if oneArtist.Id == 0 {
		Errorhandler(w, http.StatusNotFound)
		return
	}
	result := Result{
		Art: oneArtist,
		Rel: Locasya,
	}
	html, err := template.ParseFiles("templates/tracker.html")
	if err != nil {
		Errorhandler(w, http.StatusInternalServerError)
		return
	}

	err = html.Execute(w, result)
	if err != nil {
		Errorhandler(w, http.StatusInternalServerError)
		return
	}
}
