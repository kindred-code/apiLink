package photo

type Photo struct {
	Id        string `json:"id"`
	File      string `json:"file"`
	ProfileId int    `json:"profileId"`
}
