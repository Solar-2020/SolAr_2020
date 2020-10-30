package interview

import (
	"bytes"
	"encoding/json"
	"errors"
	service "github.com/Solar-2020/GoUtils/http"
	interviewApi "github.com/Solar-2020/Interview-Backend/pkg/api"
	interviewModels "github.com/Solar-2020/Interview-Backend/pkg/models"
	"github.com/Solar-2020/SolAr_Backend_2020/internal/models"
	"net/http"
)

type Client interface {
	InsertInterviews(interviews []models.Interview, postID int) (err error)
	SelectInterviews(postIDs []int) (interviews []models.Interview, err error)
	SelectInterviewsResults(postIDs []int, userID int) (interviews []interviewModels.InterviewResult, err error)
}

type client struct {
	host   string
	secret string
}

func NewClient(host string, secret string) Client {
	return &client{host: host, secret: secret}
}

type httpError struct {
	Error string `json:"error"`
}

type CreateRequest struct {
	Interviews []models.Interview `json:"interviews" validate:"required"`
	PostID     int                `json:"postID" validate:"required"`
}

func (c *client) InsertInterviews(interviews []models.Interview, postID int) (err error) {
	if len(interviews) == 0 {
		return
	}

	checkAuthRequest := CreateRequest{Interviews: interviews, PostID: postID}
	body, err := json.Marshal(checkAuthRequest)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, c.host+"/api/interview/create", bytes.NewReader(body))
	if err != nil {
		return
	}

	req.Header.Set("Authorization", c.secret)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusBadRequest:
		var httpErr httpError
		err = json.NewDecoder(resp.Body).Decode(&httpErr)
		if err != nil {
			return
		}
		return errors.New(httpErr.Error)
	default:
		return errors.New("Unexpected Server Error")
	}
}

func (c *client) SelectInterviews(postIDs []int) (interviews []models.Interview, err error) {
	endpoint := service.ServiceEndpoint{
		Service:     s,
		Endpoint:    "/api/interview",
		Method:      "POST",
		ContentType: "application/json",
	}

	message := interviewApi.GetRequest{
		Ids: postIDs,
	}

	resp := interviewApi.GetResponse{}
	err = endpoint.Send(message, &resp)
	if err != nil {
		return
	}
	interviews = FromApiInterviews(resp.Interviews)
	return
}

func (c *client) SelectInterviewsResults(postIDs []int, userID int) (interviews []interviewModels.InterviewResult, err error) {
	endpoint := service.ServiceEndpoint{
		Service:     s,
		Endpoint:    "/api/interview/list",
		Method:      "POST",
		ContentType: "application/json",
	}

	message := interviewApi.GetUniversalRequest{
		PostIDs:         postIDs,
		UserID:          userID,
		NotPassedResult: true,
	}

	resp := interviewApi.GetUniversalResponse{}
	err = endpoint.Send(message, &resp)
	if err != nil {
		return
	}
	interviews = resp.Interviews
	//interviews = FromApiInterviews(resp.Interviews)
	return
}
