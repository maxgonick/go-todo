package utils

type TodoElement struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	IsComplete  bool   `json:"isComplete"`
}

type TodoList struct {
	Elements []TodoElement `json:"Elements"`
	NextId   int           `json:"NextId"`
}
