package posts

import (
	"github.com/BarniBl/SolAr_2020/internal/models"
)

type Service interface {
	Create(request models.InputPost) (response models.Post, err error)
	GetList(request models.GetPostListRequest) (response []models.Post, err error)
}

type service struct {
	postsStorage postsStorage
	fileStorage  fileStorage
}

func NewService(postsStorage postsStorage, fileStorage fileStorage) Service {
	return &service{
		postsStorage: postsStorage,
		fileStorage:  fileStorage,
	}
}

func (s *service) Create(request models.InputPost) (response models.Post, err error) {
	if err = s.validateCreate(request); err != nil {
		return
	}

	if err = s.checkGroup(request.GroupID, request.CreateBy); err != nil {
		return
	}

	if err = s.checkFiles(request.GroupID, request.CreateBy); err != nil {
		return
	}

	if err = s.checkPhotos(request.GroupID, request.CreateBy); err != nil {
		return
	}



	response, err = s.postsStorage.InsertPost(request)
	return
}

func (s *service) validateCreate(post models.InputPost) (err error) {
	// TODO VALIDATION
	return
}

func (s *service) checkGroup(groupID, userID int) (err error) {
	return
}

func (s *service) GetList(request models.GetPostListRequest) (response []models.InputPost, err error) {

}
