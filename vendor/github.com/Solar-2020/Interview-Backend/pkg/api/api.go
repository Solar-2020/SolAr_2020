package api

import (
	"encoding/json"
	"github.com/Solar-2020/Interview-Backend/pkg/models"
)

// POST /interview/create
type CreateRequest struct {
	Interviews []models.Interview `json:"interviews" validate:"required"`
	PostID int                    `json:"postID" validate:"required"`
}

type CreateResponse struct {
	Interviews []models.Interview `json:"interviews"`
}
func (r *CreateResponse) Decode(src []byte) (err error) {
	err = json.Unmarshal(src, r)
	return
}

// POST /interview
type GetRequest struct {
	Ids []int	`json:"posts" validate:"required"`
}

type GetResponse struct {
	Interviews []models.Interview `json:"interviews"`
}
func (r *GetResponse) Decode(src []byte) (err error) {
	err = json.Unmarshal(src, r)
	return
}

// POST /interview/remove
type RemoveRequest struct {
	Ids []models.InterviewID `json:"ids" validate:"required"`
}

type RemoveResponse struct {
	Interviews []models.Interview `json:"interviews"`
}
func (r *RemoveResponse) Decode(src []byte) (err error) {
	err = json.Unmarshal(src, r)
	return
}

// GET /interview/result/:id
type ResultRequest struct {
	Id models.InterviewID `json:"id" validate:"required"`
}

type ResultResponse struct {
	models.InterviewResult
}

// POST /interview/result/:id
type SetVoteRequest struct {
	models.UserAnswer
}

type SetVoteResponse struct {
	models.InterviewResult
}
