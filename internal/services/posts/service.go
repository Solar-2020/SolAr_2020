package posts

import (
	"errors"
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

	if len(request.Files) != 0 {
		if err = s.checkFiles(request.Files, request.CreateBy); err != nil {
			return
		}
	}

	if len(request.Photos) != 0 {
		if err = s.checkPhotos(request.Photos, request.CreateBy); err != nil {
			return
		}
	}

	request.ID, err = s.postsStorage.InsertPost(request)
	return
}

func (s *service) validateCreate(post models.InputPost) (err error) {
	if len(post.Files) > 10 {
		return errors.New("В посте не может быть больше 10 файлов")
	}

	if len(post.Photos) > 10 {
		return errors.New("В посте не может быть больше 10 фотографий")
	}

	if len(post.Payments) > 10 {
		return errors.New("В посте не может быть больше 10 оплат")
	}

	if len(post.Interviews) > 10 {
		return errors.New("В посте не может быть больше 10 опросов")
	}

	return
}

func (s *service) checkGroup(groupID, userID int) (err error) {
	return
}

func (s *service) checkFiles(fileIDs []int, userID int) (err error) {
	countFiles, err := s.fileStorage.SelectCountFiles(fileIDs, userID)
	if err != nil {
		return
	}

	if countFiles != len(fileIDs) {
		return errors.New("Выбранные файлы не найдены")
	}

	return
}

func (s *service) checkPhotos(photoIDs []int, userID int) (err error) {
	countFiles, err := s.fileStorage.SelectCountPhotos(photoIDs, userID)
	if err != nil {
		return
	}

	if countFiles != len(photoIDs) {
		return errors.New("Выбранные фотографии не найдены")
	}

	return
}

func (s *service) GetList(request models.GetPostListRequest) (response []models.Post, err error) {

}
