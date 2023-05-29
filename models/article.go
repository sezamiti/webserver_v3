package models

type Article struct {
	ID      int    `json:"ID"`
	Title   string `json:"Title"`
	Author  string `json:"Author"`
	Content string `json:"Content"`
}
