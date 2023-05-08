package similarartists

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const LIMIT = 250

var API_KEY = ""

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	API_KEY = os.Getenv("API_KEY")
}

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

func ArtistGET(c *gin.Context) {
	artist := c.Param("artist")
	epGetSimilar := getEpSimilar(artist)
	client := &http.Client{}
	req, err := http.NewRequest("GET", epGetSimilar, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseJson Response
	if err := json.Unmarshal(bodyText, &responseJson); err != nil {
		c.JSON(http.StatusInternalServerError, SimilarArtistsWithError{
			Error:          err.Error(),
			Artist:         artist,
			SimilarArtists: []Artist{},
		})
	}
	sawe := SimilarArtistsWithError{
		Error:          "",
		Artist:         artist,
		SimilarArtists: []Artist{},
	}
	for _, a := range responseJson.SimilarArtists.Artist {
		match, err := strconv.ParseFloat(a.Match, 64)
		if err != nil {
			c.JSON(http.StatusInternalServerError, SimilarArtistsWithError{
				Error:          err.Error(),
				Artist:         artist,
				SimilarArtists: []Artist{},
			})
		}
		sawe.SimilarArtists = append(sawe.SimilarArtists, Artist{
			Name:  a.Name,
			Match: match,
			Url:   a.Url,
		})
	}
	c.JSON(http.StatusOK, sawe)
}
