package models

import "time"

type GetPostListRequest struct {
	userID    int
	groupID   int
	limit     int
	startFrom time.Time
}

type MainPost struct {
	ID          int         `json:"id"`
	CreateBy    int         `json:"-"`
	CreatAt     time.Time   `json:"-"`
	PublishDate time.Time   `json:"publishDate"`
	GroupID     int         `json:"groupID"`
	Text        string      `json:"text"`
	Status      string      `json:"Status"`
	Interviews  []Interview `json:"interviews"`
	Payments    []Payment   `json:"payments"`
}

type InputPost struct {
	MainPost
	Files       []WriteFile      `json:"upload"`
}

type InsertPost struct {
	MainPost
	Files       []File      `json:"upload"`
}

type Interview struct {
	ID      int      `json:"id"`
	Text    string   `json:"text"`
	Type    int      `json:"type"`
	Answers []Answer `json:"answers"`
}

type Answer struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}



type Payment struct {
	ID       int `json:"id"`
	Cost     int `json:"cost"`
	Currency int `json:"currency"`
}
