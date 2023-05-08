package similarartists

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const LIMIT = 250
const API_KEY = "b25b959554ed76058ac220b7b2e0a026"

func getEpSimilar(artist string) string {
	return strings.Join([]string{
		"http://ws.audioscrobbler.com/2.0/?method=artist.getsimilar",
		"&artist=" + artist,
		"&limit=" + strconv.Itoa(LIMIT),
		"&api_key=" + API_KEY,
		"&format=json",
	}, "")
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
	c.JSON(http.StatusOK, string(bodyText))
}
