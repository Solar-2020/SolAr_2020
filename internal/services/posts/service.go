package posts

import (
	"errors"
	"fmt"
	interviewModels "github.com/Solar-2020/Interview-Backend/pkg/models"
	"github.com/Solar-2020/SolAr_Backend_2020/internal/clients/account"
	"github.com/Solar-2020/SolAr_Backend_2020/internal/clients/group"
	"github.com/Solar-2020/SolAr_Backend_2020/internal/models"
	"sort"
)

type Service interface {
	Create(request models.InputPost) (response models.Post, err error)
	GetList(request models.GetPostListRequest) (response []models.PostResult, err error)
}

type service struct {
	postsStorage     postStorage
	uploadStorage    uploadStorage
	interviewStorage interviewStorage
	paymentStorage   paymentStorage
	groupClient      group.Client
	accountClient    account.Client
}

func NewService(postsStorage postStorage, uploadStorage uploadStorage, interviewStorage interviewStorage, paymentStorage paymentStorage, groupClient group.Client, accountClient account.Client) Service {
	return &service{
		postsStorage:     postsStorage,
		uploadStorage:    uploadStorage,
		interviewStorage: interviewStorage,
		paymentStorage:   paymentStorage,
		groupClient:      groupClient,
		accountClient:    accountClient,
	}
}

func (s *service) Create(request models.InputPost) (response models.Post, err error) {
	if err = s.validateCreate(request); err != nil {
		return
	}

	roleID, err := s.groupClient.GetUserRole(request.CreateBy, request.GroupID)
	if err != nil {
		err = fmt.Errorf("restricted")
		return
	}

	if roleID > 2 {
		return response, errors.New("permission denied")
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

func (s *service) GetList(request models.GetPostListRequest) (response []models.PostResult, err error) {
	roleID, err := s.groupClient.GetUserRole(request.UserID, request.GroupID)
	if err != nil {
		err = fmt.Errorf("restricted")
		return
	}

	if roleID > 3 {
		return response, errors.New("permission denied")
	}

	posts, err := s.postsStorage.SelectPosts(request)
	if err != nil {
		return
	}

	if len(posts) == 0 {
		return
	}

	postsMap := make(map[int]models.PostResult)
	for index, post := range posts {
		postsMap[post.ID] = models.PostResult{
			ID:          post.ID,
			CreateBy:    post.CreateBy,
			CreatAt:     post.CreatAt,
			PublishDate: post.PublishDate,
			GroupID:     post.GroupID,
			Text:        post.Text,
			Status:      post.Status,
			Photos:      make([]models.Photo, 0),
			Files:       make([]models.File, 0),
			Interviews:  make([]interviewModels.InterviewResult, 0),
			Payments:    make([]models.Payment, 0),
			Order:       index,
		}
	}

	postIDs := make([]int, 0)
	for i, _ := range posts {
		postIDs = append(postIDs, posts[i].ID)
	}

	interviews, err := s.interviewStorage.SelectInterviewsResults(postIDs, request.UserID)
	if err != nil {
		return
	}

	payments, err := s.paymentStorage.SelectPayments(postIDs)
	if err != nil {
		return
	}

	matchPostPhoto, err := s.postsStorage.SelectPhotoIDs(postIDs)
	if err != nil {
		return
	}

	photoIDs := make([]int, 0)
	for i, _ := range matchPostPhoto {
		photoIDs = append(photoIDs, matchPostPhoto[i].PhotoID)
	}

	matchPostFile, err := s.postsStorage.SelectFileIDs(postIDs)
	if err != nil {
		return
	}

	fileIDs := make([]int, 0)
	for i, _ := range matchPostFile {
		fileIDs = append(fileIDs, matchPostFile[i].FileID)
	}

	photos, err := s.uploadStorage.SelectPhotos(photoIDs)
	if err != nil {
		return
	}

	files, err := s.uploadStorage.SelectFiles(fileIDs)
	if err != nil {
		return
	}

	for _, interview := range interviews {
		tempPost := postsMap[interview.PostID]
		tempPost.Interviews = append(tempPost.Interviews, interview)
		postsMap[interview.PostID] = tempPost
	}

	for _, payment := range payments {
		tempPost := postsMap[payment.PostID]
		tempPost.Payments = append(tempPost.Payments, payment)
		postsMap[payment.PostID] = tempPost
	}

	for _, match := range matchPostPhoto {
		tempPost := postsMap[match.PostID]
		tempPost.Photos = append(tempPost.Photos, photos[match.PhotoID])
		postsMap[match.PostID] = tempPost
	}

	for _, match := range matchPostFile {
		tempPost := postsMap[match.PostID]
		tempPost.Files = append(tempPost.Files, files[match.FileID])
		postsMap[match.PostID] = tempPost
	}

	for _, post := range postsMap {
		response = append(response, post)
	}

	sortPost := models.Posts{Posts: response}
	sort.Sort(&sortPost)

	for i, _ := range sortPost.Posts {
		var user models.User
		user, err = s.accountClient.GetUserByID(sortPost.Posts[i].CreateBy)
		if err != nil {
			return
		}
		sortPost.Posts[i].Author = user
	}

	return sortPost.Posts, nil
}
