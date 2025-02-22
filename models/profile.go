package models

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
