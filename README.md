Hi, this is the first program I've ever written in Go. Feedback is welcome, and expect strange code. This is essentially a proof of concept, I will be implementing more API interactions later, and I'll be finding a better way to actually publish the functionality, I'm not a huge fan of the web server setup. **This also might be against the Spotify TOS to actually use...**

# spotify-api-tool-go

This program uses the Spotify Web Player API to recreate missing functionality from the Spotify Developer API. Specifically, this allows you to **retrieve user followers**, which has been a largely requested feature among the Spotify developer community. The program itself is a locally hosted web server with endpoints to interact with the Spotify API. I'll work on a different implementation of the actual functionality later.

### Usage

- `go run .` - run the program
- `go build` - build the program

### API Usage

The web server is hosted on port `8765` by default.

---
- `GET /auth`

Returns `{client_token: {...}, access_token: {...}}`

Modeled after:

```go
type AccessTokenObject struct {
	AccessToken string `json:"accessToken"`
	TokenType   string `json:"tokenType"`
	ExpiresIn   int    `json:"expiresIn"`
	ClientId    string `json:"clientId"`
}

type ClientTokenObject struct {
	ResponseType string `json:"response_type"`
	GrantedToken struct {
		Token               string   `json:"token"`
		ExpiresAfterSeconds int      `json:"expires_after_seconds"`
		RefreshAfterSeconds int      `json:"refresh_after_seconds"`
		Domains             []Domain `json:"domains"`
	} `json:"granted_token"`
}
```
---
- `GET /profile?username`

**Requires** the `access-token` and `client-token` headers for their respective token strings. The `username` query parameter takes the Spotify username. This is different from the common display name; it is found inside of a Spotify profile URL `https://open.spotify.com/user/ofngctk005r8jdcm2r02176cc` where `ofngctk005r8jdcm2r02176cc` is the username. This is basically Spotify's user ID.

Returns `{ uri: "...", ... }`

Modeled after:

```go
type Profile struct {
	URI                       string      `json:"uri"`
	Name                      string      `json:"name"`
	ImageURL                  *string     `json:"image_url"`
	FollowersCount            *int        `json:"followers_count"`
	FollowingCount            *int        `json:"following_count"`
	RecentlyPlayedArtists     *[]Artist   `json:"recently_played_artists"`
	PublicPlaylists           *[]Playlist `json:"public_playlists"`
	TotalPublicPlaylistsCount *int        `json:"total_public_playlists_count"`
	HasSpotifyName            bool        `json:"has_spotify_name"`
	Color                     int         `json:"color"`
	AllowFollows              bool        `json:"allow_follows"`
	ShowFollows               bool        `json:"show_follows"`
}
```
---
- `GET /followers?username`

**Requires** the `access-token` and `client-token` headers for their respective token strings. The `username` query parameter takes the Spotify username.

Returns `{ "profiles": [{ "uri": "..." }, {"uri": "..."}] }`

Modeled after:

```go
type Follower struct {
	URI            string  `json:"uri"`
	Name           string  `json:"name"`
	ImageURL       *string `json:"image_url"`
	FollowersCount *int    `json:"followers_count"`
	Color          int     `json:"color"`
}

type FollowersResponse struct {
	Followers []Follower `json:"profiles"`
}
```

