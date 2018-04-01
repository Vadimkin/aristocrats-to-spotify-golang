package spotty

import (
    "github.com/zmb3/spotify"
    "log"
    "fmt"
)

func SearchTrack(
    authorName string,
    trackName string,
) (spotify.FullTrack, bool) {
    var track spotify.FullTrack
    searchQuery := authorName + " " + trackName
    results, err := client.Search(searchQuery, spotify.SearchTypeTrack)
    
    if err != nil {
        log.Fatal(err)
        return track, true
    }
    
    isNotFound := results.Tracks == nil
    if results.Tracks != nil && len(results.Tracks.Tracks) > 0 {
        track = results.Tracks.Tracks[0]
    } else {
        isNotFound = true
    }
    
    return track, isNotFound
}

func addTrackToPlaylist(playlistID spotify.ID, track spotify.FullTrack) {
    res, err := client.AddTracksToPlaylist(
        SPOTIFY_USER_ID,
        playlistID,
        track.ID,
    )
    fmt.Println(track.URI)
    fmt.Println(res)
    fmt.Println(err)
}

func filterOnlyNewTracks(playlistID spotify.ID, tracks []spotify.FullTrack) []spotify.FullTrack {
    var tracksMap = map[spotify.ID]spotify.FullTrack{}
    
    playlistTracks, err := client.GetPlaylistTracks(
        SPOTIFY_USER_ID,
        playlistID,
    )
    
    if err != nil {
        log.Fatal(err)
    }
    
    
    for _, track := range tracks {
        tracksMap[track.ID] = track
    }
    
    // Didn't found any set operations in Go
    for _, track := range playlistTracks.Tracks {
        delete(tracksMap, track.Track.ID)
    }
    
    
    // Convert back map to list
    tracks = make([]spotify.FullTrack, 0, len(tracksMap))
    for _, track := range tracksMap {
        tracks = append(tracks, track)
    }
    
    return tracks
}

func AddTracksToPlaylist(tracks []spotify.FullTrack) {
    var playlistID = spotify.ID(SPOTIFY_PLAYLIST_ID)
    
    tracks = filterOnlyNewTracks(playlistID, tracks)
    
    for _, track := range tracks {
       addTrackToPlaylist(playlistID, track)
    }
}