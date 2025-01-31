package cmd

type todoElement struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	IsComplete  bool   `json:"isComplete"`
}
