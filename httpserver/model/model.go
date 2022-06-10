package model

// todo item
type todoItem struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

// albums slice to seed record album data.
var Todos = []todoItem{
	{ID: "1", Text: "Eat", Completed: true},
	{ID: "2", Text: "Surf", Completed: false},
	{ID: "3", Text: "Sleep", Completed: false},
	{ID: "4", Text: "Repeat", Completed: true},
}
