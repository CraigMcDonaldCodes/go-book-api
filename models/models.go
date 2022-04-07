package models

import "time"

type Book struct {
	Id          int64     `json:"id"`
	Isbn        string    `json:"isbn"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Publisher   string    `json:"publisher"`
	PublishDate time.Time `json:"publishdate"`
}
