package models

type Playlist struct {
	URI            string `json:"uri"`
	Name           string `json:"name"`
	ImageURL       string `json:"image_url"`
	FollowersCount *int   `json:"followers_count"`
	OwnerName      string `json:"owner_name"`
	OwnerURI       string `json:"owner_uri"`
}
