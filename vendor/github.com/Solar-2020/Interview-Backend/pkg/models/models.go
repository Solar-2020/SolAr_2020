package models

type InterviewID int
type AnswerID int
type VoteID int
type InterviewType  int

type InterviewFrame struct {
	ID     InterviewID    `json:"id"`
	Text   string `json:"text"`
	Type   InterviewType    `json:"type"`
	PostID int    `json:"postID"`
	Status int    `json:"status"` //Проголосовал юзер или нет
}

type Interview struct {
	InterviewFrame
	Answers []Answer `json:"answers"`
}

type Answer struct {
	ID          AnswerID    `json:"id"`
	Text        string `json:"text"`
	InterviewID InterviewID    `json:"interviewID"`
}

type InterviewResult struct {
	InterviewFrame
	Answers []AnswerResult `json:"answers"`
}

type AnswerResult struct {
	Answer
	AnswerCount int `json:"answerCount"`
}

type UserAnswer struct {
	ID          VoteID `json:"id"`
	InterviewID InterviewID `json:"interviewID" validate:"required"`
}

type UserAnswers struct {
	PostID      int   `json:"postID"`
	InterviewID InterviewID   `json:"interviewID"`
	UserID      int   `json:"-"`
	AnswerIDs   []AnswerID `json:"answers"`
}
