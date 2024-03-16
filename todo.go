package todo

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Describtion string `json:"Describtion"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Describtion string `json:"Describtion"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	Id     int
	UserId int
	ListId int
}
