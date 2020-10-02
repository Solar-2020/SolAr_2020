package models

import "time"

type GetPostListRequest struct {
	UserID    int
	GroupID   int
	Limit     int
	StartFrom time.Time
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
	ID          int         `json:"id"`
	CreateBy    int         `json:"-"`
	CreatAt     time.Time   `json:"-"`
	PublishDate time.Time   `json:"publishDate"`
	GroupID     int         `json:"groupID"`
	Text        string      `json:"text"`
	Status      string      `json:"Status"`
	Photos      []int       `json:"photos"`
	Files       []int       `json:"files"`
	Interviews  []Interview `json:"interviews"`
	Payments    []Payment   `json:"payments"`
}

type Post struct {
	ID          int         `json:"id"`
	CreateBy    int         `json:"-"`
	CreatAt     time.Time   `json:"-"`
	PublishDate time.Time   `json:"publishDate"`
	GroupID     int         `json:"groupID"`
	Text        string      `json:"text"`
	Status      string      `json:"Status"`
	Photos      []Photo     `json:"photos"`
	Files       []File      `json:"files"`
	Interviews  []Interview `json:"interviews"`
	Payments    []Payment   `json:"payments"`
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