package todo

type TodoList struct {
	Id         int    `json:"id" bd:"id"`
	Title      string `json:"title" bd:"title" binding:"required"`
	Description string `json:"description" bd:"description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}
