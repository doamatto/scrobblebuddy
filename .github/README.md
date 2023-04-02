Scrobble Buddy is a simple tool to let you scrobble with other Last.fm users.

Usage: `scrobblebuddy <last_fm_user_to_scrobble_along_with>`

Configuration options are within `.config/scrobblebuddy.json` (or wherever `XDG_CONFIG_HOME` is set to) and are follows:
  - `username` (string): your Last.fm username
  - `password` (string): your Last.fm password (this should be )
  - `server_url` (URL): the URL for where you scrobble to (setting this value will override the default, which is Last.fm's server: `http://ws.audioscrobbler.com/2.0/`)
  - `api_key` (string) and `api_secret`: your Last.fm credentials (you can [get an API key and secret here](https://www.last.fm/api/account/create))