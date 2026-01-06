package entitys

import "errors"

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

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Complete    *bool   `json:"complete"`
}

func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Complete == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
