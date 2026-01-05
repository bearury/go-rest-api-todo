package entitys

type Todo struct {
	Id          string `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type UserList struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	ListId string `json:"list_id"`
}

type TodoItem struct {
	Id          string `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Complete    bool   `json:"complete" default:"false"`
}

type ListItem struct {
	Id     string `json:"id"`
	ListId string `json:"list_id"`
	ItemId string `json:"item_id"`
}
