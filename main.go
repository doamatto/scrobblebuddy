package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/shkh/lastfm-go/lastfm"
)

const Version string = "v0.1.2"

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	APIKey string `json:"api_key"`
	APISecret string `json:"api_secret"`
}

func main() {
	log.Printf("[SCROBBLEBUDDY] You are using %s.\n", Version)

	// Get user to listen along with
	if len(os.Args) < 2 {
		log.Fatalln("[SCROBBLEBUDDY] You need to specify a user to scrobble alongside. The command should look similar to the below, replacing `bbcradio1` with the user you want to scrobble with.\n\nscrobblebuddy bbcradio1")
	}
	user := os.Args[1]

	// Load configuration file
	var confHome string = os.Getenv("XDG_CONFIG_HOME")
	if confHome == "" {
		confHome = fmt.Sprintf("%s/.config", os.Getenv("HOME"))
	}
	confFileLoc := fmt.Sprintf("%s/scrobblebuddy.json", confHome)

	conf, err := os.ReadFile(confFileLoc)
	if err != nil {
		log.Fatalf("[SCROBBLEBUDDY] Failed to load configuration file %s. Error to follow:\n\n%s", confFileLoc, err)
	}
	var config Config
	if err = json.Unmarshal(conf, &config); err != nil {
		log.Fatalf("[SCROBBLEBUDDY] Failed to read configuration file %s. Error to follow:\n\n%s", confFileLoc, err)
	}

	// Authenticate with Last.fm
	api := lastfm.New(config.APIKey, config.APISecret)
	err = api.Login(config.Username, config.Password)
	if err != nil {
		log.Fatalf("[SCROBBLEBUDDY] Failed to authenticate as %s. Error to follow:\n\n%s", config.Username, err)
	}
	log.Printf("[SCROBBLEBUDDY] Authenticated as %s.\n", config.Username)

	// Start loop for listen along
	err = scrobblealong(api, user, config)
	if err != nil {
		log.Fatal(err)
	}
}

// We use a seperate function so that we can
// loop the function without having to duplicate a lot of code.
func scrobblealong(api *lastfm.Api, user string, conf Config) (err error) {
	// Check if user is listening to music (must be "Playing now")
	data, err := api.User.GetRecentTracks(lastfm.P{
			"user": user,
			"limit": 1, // we don't need the default 50
	})
	if err != nil { return err }

	// Check if song is currently playing
	if data.Tracks[0].NowPlaying == "true" {
		// Check if song is/was being scrobbled now
		curTrack, err := api.User.GetRecentTracks(lastfm.P{
			"user": conf.Username,
			"limit": 1, // we don't need the default 50
		})
		if err != nil { return err }
		if data.Tracks[0].Name == curTrack.Tracks[0].Name {
			// Timeout for 30 seconds
			time.Sleep(30 * time.Second)

			// Rinse and repeat
			err := scrobblealong(api, user, conf)
			if err != nil {	return err }	
		}

		log.Printf("Scrobbling %s by %s...", data.Tracks[0].Name, data.Tracks[0].Artist.Name)
		// Scrobble the song
		_, err = api.Track.UpdateNowPlaying(lastfm.P{
			"artist": data.Tracks[0].Artist.Name,
			"track": data.Tracks[0].Name,
		})
		if err != nil { return err }
		_, err = api.Track.Scrobble(lastfm.P{
			"artist": data.Tracks[0].Artist.Name,
			"track": data.Tracks[0].Name,
			"timestamp": time.Now().Unix(),
		})
		if err != nil { return err }

		// Timeout for 30 seconds
		time.Sleep(30 * time.Second)

		// Rinse and repeat
		err = scrobblealong(api, user, conf)
		if err != nil {	return err }
	}
	return nil // this should never happen
}
