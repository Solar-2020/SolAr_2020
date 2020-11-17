package posts

import (
	interviewModels "github.com/Solar-2020/Interview-Backend/pkg/models"
	"github.com/Solar-2020/SolAr_Backend_2020/internal/models"
)

const (
	GetPostActionID    = 4
	CreatePostActionID = 5
	EditPostActionID   = 6
	DeletePostActionID = 7
	MarkPostActionID   = 8

	PostStatusCreating = 1
	PostStatusCreated  = 2
	PostStatusRemoved  = 3
)

type postStorage interface {
	InsertPost(inputPost models.InputPost) (postID int, err error)
	UpdatePostStatus(postID int, groupID int, status int) (err error)
	SetMark(postID int, mark bool, group int) (err error)

	SelectFileIDs(postIDs []int) (matches []models.PostFileMatch, err error)
	SelectPhotoIDs(postIDs []int) (matches []models.PostPhotoMatch, err error)

	SelectPosts(request models.GetPostListRequest) (posts []models.InputPost, err error)

	SelectInterviews(postIDs []int) (interviews []models.Interview, err error)
}

type uploadStorage interface {
	SelectCountFiles(fileIDs []int, userID int) (countFiles int, err error)
	SelectCountPhotos(photoIDs []int, userID int) (countFiles int, err error)

	SelectFiles(fileIDs []int) (files map[int]models.File, err error)
	SelectPhotos(photoIDs []int) (photos map[int]models.Photo, err error)
}

type interviewStorage interface {
	InsertInterviews(interviews []models.Interview, postID int) (err error)

	SelectInterviewsResults(postIDs []int, userID int) (interviews []interviewModels.InterviewResult, err error)
}

type paymentClient interface {
	Create(createRequest models.CreateRequest) (createdPayments []models.Payment, err error)
	GetByPostIDs(postIDs []int) (payments []models.Payment, err error)
}

type groupClient interface {
	CheckPermission(userID, groupId, actionID int) (err error)
}
