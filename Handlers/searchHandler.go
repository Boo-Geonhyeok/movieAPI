package handlers

import (
	"encoding/json"
	"net/http"

	fetch "github.com/movie/Fetch"
	models "github.com/movie/Models"
)

func GetSearchedMovieList(w http.ResponseWriter, r *http.Request) {
	type Title struct {
		Title string
	}
	search := Title{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)

		return
	}
	response, err := fetch.FetchTitle(search.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)

		return
	}
	result := models.TitleResult{}
	json.Unmarshal(response, &result)
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Location", "http://127.0.0.1:8080/?search="+search.Title)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result.Results)
}
