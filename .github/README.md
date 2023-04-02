Scrobble Buddy is a simple tool to let you scrobble with other Last.fm users.

## Building
0. Install [Golang](https://golang.org/dl)
1. Clone this repository (`git clone https://github.com/doamatto/scrobblebuddy.git`)
2. Build this app (`go build -o scrobblebuddy main.go`)

You can now use `scrobblebuddy`

## Usage
Usage: `scrobblebuddy <last_fm_user_to_scrobble_along_with>`

Configuration options are within `.config/scrobblebuddy.json` (or wherever `XDG_CONFIG_HOME` is set to) and are follows:
  - `username` (string): your Last.fm username
  - `password` (string): your Last.fm password
  - `api_key` (string) and `api_secret` (string): your Last.fm credentials (you can [get an API key and secret here](https://www.last.fm/api/account/create))

## Acknowledgements
This program is licensed under the 3-Clause BSD license. A copy of this license is in the `LICENSE` file in the root of this repository.

This program uses [Kohei Shitaune](https://github.com/shkh)'s [lastfm-go library](https://github.com/shkh/lastfm-go/blob/89a801c244e0e5c320fcd4416e7a30520dc8a233/lastfm/lastfm.go) (MIT). This program was inspiried by [Amr Hassan](https://github.com/amrhassan)'s [bbcscrobbler program](https://github.com/amrhassan/bbcscrobbler).
