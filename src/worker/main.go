package worker

import (
	"github.com/anaskhan96/soup"
	"github.com/levigross/grequests"
	"log"
	"strings"
)

func getPage(page string) string {
	resp, err := grequests.Get(page, nil)

	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	return resp.String()
}

func getTrackList(content string) []Track {
	var tracks []Track
	soupDocument := soup.HTMLParse(content)

	links := soupDocument.FindAll("div", "class", "track")

	for _, link := range links {
		trackName := link.Text()
		trackName = strings.TrimLeft(trackName, " - ")

		trackAuthor := link.Find("span", "class", "artist").Text()

		track := Track{trackName, trackAuthor}
		tracks = append(tracks, track)
	}

	return tracks
}

func main() {
	response := getPage(AristocratsMain)
	tracks := getTrackList(response)

	for _, track := range tracks {
		log.Println(track.author, "-", track.title)
	}
}
