package models

type Artist struct {
	URI            string `json:"uri"`
	Name           string `json:"name"`
	ImageURL       string `json:"image_url"`
	FollowersCount *int   `json:"followers_count"`
}
