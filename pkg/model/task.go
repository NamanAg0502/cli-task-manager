package model

type Task struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
