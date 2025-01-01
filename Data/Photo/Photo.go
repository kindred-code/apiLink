package photo

type Photo struct {
	Id     string `json:"id"`
	File   string `form:"file"`
	UserId int    `json:"userId"`
}
