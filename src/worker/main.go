package main

import (
    "github.com/zmb3/spotify"
    "worker/spotty"
    "fmt"
)

func main() {
    var tracks []spotify.FullTrack
    //response := parser.GetPage(parser.AristocratsMain)
    //var tracks = parser.GetTrackList(response)
    //
    //
    //for _, track := range tracks {
    //    log.Println(track.Author, '-', track.Title)
    //}
    
    track, isNotFound := spotty.SearchTrack("Hurts", "Ready To Go")
    fmt.Println(isNotFound)
    fmt.Println(track.ID)
    
    tracks = append(tracks, track)
    
    spotty.AddTracksToPlaylist(tracks)
}

