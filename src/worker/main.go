package main

import (
    "log"
    "worker/parser"
)

func main() {
    response := parser.GetPage(parser.AristocratsMain)
    var tracks = parser.GetTrackList(response)
    

    for _, track := range tracks {
        log.Println(track.Author, '-', track.Title)
    }
}

