package posts

import (
	interviewModels "github.com/Solar-2020/Interview-Backend/pkg/models"
	"github.com/Solar-2020/SolAr_Backend_2020/internal/models"
)

type postStorage interface {
	InsertPost(inputPost models.InputPost) (postID int, err error)

	UpdatePostStatus(postID int, status int) (err error)

	SelectFileIDs(postIDs []int) (matches []models.PostFileMatch, err error)
	SelectPhotoIDs(postIDs []int) (matches []models.PostPhotoMatch, err error)

	SelectPosts(request models.GetPostListRequest) (posts []models.InputPost, err error)

	SelectPayments(postIDs []int) (payments []models.Payment, err error)

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
	SelectInterviews(postIDs []int) (interviews []models.Interview, err error)

	SelectInterviewsResults(postIDs []int, userID int) (interviews []interviewModels.InterviewResult, err error)
}

type paymentStorage interface {
	InsertPayments(payments []models.Payment, postID int) (err error)
	SelectPayments(postIDs []int) (payments []models.Payment, err error)
}
