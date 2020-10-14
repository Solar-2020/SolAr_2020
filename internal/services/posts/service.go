package posts

import (
	"errors"
	"github.com/Solar-2020/SolAr_Backend_2020/internal/models"
)

type Service interface {
	Create(request models.InputPost) (response models.Post, err error)
	GetList(request models.GetPostListRequest) (response []models.InputPost, err error)
}

type service struct {
	postsStorage     postStorage
	uploadStorage    uploadStorage
	interviewStorage interviewStorage
	paymentStorage   paymentStorage
}

func NewService(postsStorage postStorage, uploadStorage uploadStorage, interviewStorage interviewStorage, paymentStorage paymentStorage) Service {
	return &service{
		postsStorage:     postsStorage,
		uploadStorage:    uploadStorage,
		interviewStorage: interviewStorage,
		paymentStorage:   paymentStorage,
	}
}

func (s *service) Create(request models.InputPost) (response models.Post, err error) {
	if err = s.validateCreate(request); err != nil {
		return
	}

	if err = s.checkGroup(request.GroupID, request.CreateBy); err != nil {
		return
	}

	if err = s.checkFiles(request.Files, request.CreateBy); err != nil {
		return
	}

	if err = s.checkPhotos(request.Photos, request.CreateBy); err != nil {
		return
	}

	response.ID, err = s.postsStorage.InsertPost(request)
	if err != nil {
		return
	}

	err = s.interviewStorage.InsertInterviews(request.Interviews, response.ID)
	if err != nil {
		return
	}

	err = s.paymentStorage.InsertPayments(request.Payments, response.ID)
	if err != nil {
		return
	}

	// TODO CHANGE TO CONST
	err = s.postsStorage.UpdatePostStatus(response.ID, 2)
	if err != nil {
		return
	}

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
	countFiles, err := s.uploadStorage.SelectCountFiles(fileIDs, userID)
	if err != nil {
		return
	}

	if countFiles != len(fileIDs) {
		return errors.New("Выбранные файлы не найдены")
	}

	return
}

func (s *service) checkPhotos(photoIDs []int, userID int) (err error) {
	countFiles, err := s.uploadStorage.SelectCountPhotos(photoIDs, userID)
	if err != nil {
		return
	}

	if countFiles != len(photoIDs) {
		return errors.New("Выбранные фотографии не найдены")
	}

	return
}

func (s *service) GetList(request models.GetPostListRequest) (response []models.InputPost, err error) {
	posts, err := s.postsStorage.SelectPosts(request)
	if err != nil {
		return
	}

	postIDs := make([]int, 0)
	for i, _ := range posts {
		postIDs = append(postIDs, posts[i].ID)
	}

	response = posts

	//interview, err := s.postStorage.SelectInterviews(postIDs)
	//if err != nil {
	//	return
	//}
	//
	//payments, err := s.postStorage.SelectPayments(postIDs)
	//if err != nil {
	//	return
	//}
	//
	//files, err := s.uploadStorage.SelectFiles(postIDs)
	//if err != nil {
	//	return
	//}
	//
	//photos, err := s.uploadStorage.SelectPhotos(postIDs)
	//if err != nil {
	//	return
	//}
	//
	//fmt.Println(posts, interview, payments, files, photos)
	return
}
