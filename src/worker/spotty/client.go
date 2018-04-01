package spotty

import (
    "github.com/zmb3/spotify"
    "golang.org/x/oauth2/clientcredentials"
    "log"
    "context"
)

func getClient() spotify.Client {
    config := &clientcredentials.Config{
        ClientID:     SPOTIFY_ID,
        ClientSecret: SPOTIFY_SECRET,
        TokenURL:     spotify.TokenURL,
    }
    token, err := config.Token(context.Background())
    if err != nil {
        log.Fatalf("couldn't get token: %v", err)
    }
    
    client := spotify.Authenticator{}.NewClient(token)
    
    return client
}

var client = getClient()