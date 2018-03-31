package main

import (
	"github.com/levigross/grequests"
	"log"
	"github.com/anaskhan96/soup"
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
	var soupDocument = soup.HTMLParse(content)

	links := soupDocument.FindAll("div", "class", "track")

	for _, link := range links {
		var trackName = link.Text()
		trackName = strings.TrimLeft(trackName, " - ")

		var trackAuthor = link.Find("span", "class", "artist").Text()

		var track = Track{trackName, trackAuthor}
		tracks = append(tracks, track)
	}

	return tracks
}

func main() {
	var response = getPage(AristocratsMain)
	var tracks = getTrackList(response)

	for _, track := range tracks {
		log.Println(track.author, "-", track.title)
	}
}
