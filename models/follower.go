package models

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
