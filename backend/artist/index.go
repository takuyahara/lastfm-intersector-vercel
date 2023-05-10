package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const LIMIT = 250

var API_KEY = os.Getenv("API_KEY")

func getEpSimilar(artist string) string {
	return strings.Join([]string{
		"http://ws.audioscrobbler.com/2.0/?method=artist.getsimilar",
		"&artist=" + artist,
		"&limit=" + strconv.Itoa(LIMIT),
		"&api_key=" + API_KEY,
		"&format=json",
	}, "")
}

type Response struct {
	SimilarArtists struct {
		Artist []ArtistRaw `json:"artist"`
	} `json:"similarartists"`
}
type ArtistRaw struct {
	Name  string `json:"name"`
	Match string `json:"match"`
	Url   string `json:"url"`
}
type Artist struct {
	Name  string  `json:"name"`
	Match float64 `json:"match"`
	Url   string  `json:"url"`
}
type SimilarArtistsWithError struct {
	Error          string   `json:"error"`
	Artist         string   `json:"artist"`
	SimilarArtists []Artist `json:"similarartists"`
}

func getArtist(artist string) SimilarArtistsWithError {
	epGetSimilar := getEpSimilar(artist)
	client := &http.Client{}
	req, err := http.NewRequest("GET", epGetSimilar, nil)
	if err != nil {
		return SimilarArtistsWithError{
			Error:          err.Error(),
			Artist:         artist,
			SimilarArtists: []Artist{},
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return SimilarArtistsWithError{
			Error:          err.Error(),
			Artist:         artist,
			SimilarArtists: []Artist{},
		}
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return SimilarArtistsWithError{
			Error:          err.Error(),
			Artist:         artist,
			SimilarArtists: []Artist{},
		}
	}
	var responseJson Response
	if err := json.Unmarshal(bodyText, &responseJson); err != nil {
		return SimilarArtistsWithError{
			Error:          err.Error(),
			Artist:         artist,
			SimilarArtists: []Artist{},
		}
	}
	sawe := SimilarArtistsWithError{
		Error:          "",
		Artist:         artist,
		SimilarArtists: []Artist{},
	}
	for _, a := range responseJson.SimilarArtists.Artist {
		match, err := strconv.ParseFloat(a.Match, 64)
		if err != nil {
			return SimilarArtistsWithError{
				Error:          err.Error(),
				Artist:         artist,
				SimilarArtists: []Artist{},
			}
		}
		sawe.SimilarArtists = append(sawe.SimilarArtists, Artist{
			Name:  a.Name,
			Match: match,
			Url:   a.Url,
		})
	}
	return sawe
}

func Handler(w http.ResponseWriter, r *http.Request) {
	artist := r.URL.Query().Get("artist")
	similarartists := getArtist(artist)
	bytes, _ := json.Marshal(similarartists)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,OPTIONS,PATCH,DELETE,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "X-CSRF-Token, X-Requested-With, Accept, Accept-Version, Content-Length, Content-MD5, Content-Type, Date, X-Api-Version")
	fmt.Fprintf(w, string(bytes))
}
