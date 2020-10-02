package upload

import (
	"github.com/BarniBl/SolAr_2020/internal/models"
)

type uploadStorage interface {
	SaveFile(file models.WriteFile) (fileView models.File, err error)
	SavePhoto(file models.WritePhoto) (fileView models.Photo, err error)
	//SaveFiles(files []models.WriteFile) (fileView []models.File, err error)

	InsertFile(file models.File, userID int) (fileID int, err error)
	InsertPhoto(photo models.Photo, userID int) (photoID int, err error)
}
