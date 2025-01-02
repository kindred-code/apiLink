package profile

type Auth struct {
	ProfileId int    `json:"profileId"`
	Token     string `json:"token"`
}
