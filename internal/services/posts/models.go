package posts

import (
	"github.com/BarniBl/SolAr_2020/internal/models"
)

type postsStorage interface {
	InsertPost(inputPost models.InputPost) (postID int, err error)
	SelectPosts(request models.GetPostListRequest) (posts []models.InputPost, err error)
}

type uploadStorage interface {
	SelectCountFiles(fileIDs []int, userID int) (countFiles int, err error)
	SelectCountPhotos(photoIDs []int, userID int) (countFiles int, err error)

	SelectFiles(fileIDs []int) (files []models.File, err error)
	SelectPhotos(photoIDs []int) (photos []models.Photo, err error)
}
