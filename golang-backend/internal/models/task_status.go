package models

type TaskStatus struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Index int    `json:"index"`
	Color string `json:"color"`
}
