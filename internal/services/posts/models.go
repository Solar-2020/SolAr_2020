package posts

import (
	"github.com/BarniBl/SolAr_2020/internal/models"
)

type postsStorage interface {
	InsertPost(models.Post) (post models.Post, err error)
}

type fileStorage interface {
	SaveFile(file models.WriteFile) (fileView models.File, err error)
	SaveFiles(files []models.WriteFile) (fileView []models.File, err error)
}