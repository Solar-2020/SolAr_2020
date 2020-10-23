package interviewStorage

import (
	service "github.com/Solar-2020/GoUtils/http"
	interviewApi "github.com/Solar-2020/Interview-Backend/pkg/api"
	"github.com/Solar-2020/SolAr_Backend_2020/internal/models"
)

type ServiceStorage struct {
	service.Service
	address string
}

func (s *ServiceStorage) Address() string {
	return s.address
}

func NewStorageProxy(serviceAddress string) Storage {
	return &ServiceStorage{
		address: serviceAddress,
	}
}

func (s *ServiceStorage) InsertInterviews(interviews []models.Interview, postID int) (err error) {
	endpoint := service.ServiceEndpoint{
		Service:   	 s,
		Endpoint:    "/interview/create",
		Method:      "POST",
		ContentType: "application/json",
	}

	message := interviewApi.CreateRequest{
		Interviews: ToApiInterviews(interviews),
		PostID:     postID,
	}

	resp := interviewApi.CreateResponse{}
	err = endpoint.Send(message, &resp)
	if err != nil {
		return
	}
	interviews = FromApiInterviews(resp.Interviews)
	return err

}

func (s *ServiceStorage) SelectInterviews(postIDs []int) (interviews []models.Interview, err error) {
	endpoint := service.ServiceEndpoint{
		Service:   	 s,
		Endpoint:    "/interview",
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
