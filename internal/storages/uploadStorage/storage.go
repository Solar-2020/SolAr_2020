package uploadStorage

import (
	"crypto/md5"
	"database/sql"
	"errors"
	"github.com/BarniBl/SolAr_2020/internal/models"
	"os"
	"strings"
	"time"
)

type Storage interface {
	SaveFile(file models.WriteFile) (fileView models.File, err error)
	SavePhoto(file models.WritePhoto) (fileView models.Photo, err error)
	//SaveFiles(files []models.WriteFile) (fileView []models.File, err error)

	InsertFile(file models.File) (fileID int, err error)
	InsertPhoto(photo models.Photo) (photoID int, err error)
}

type storage struct {
	photoPath string
	filePath  string
	db        *sql.DB
}

func NewStorage(photoPath, filePath string, db *sql.DB) Storage {
	return &storage{
		photoPath: photoPath,
		filePath:  filePath,
		db:        db,
	}
}

func (s *storage) SaveFile(file models.WriteFile) (fileView models.File, err error) {
	fileName, err := s.createFileName(file.Name)
	if err != nil {
		return
	}

	filePath, err := s.createFilePath(file.Name)
	if err != nil {
		return
	}

	fileView.Name = fileName
	fileView.URL = s.filePath + "/" + filePath + "/" + fileName

	writeFile, err := os.Create(fileView.URL)
	if err != nil {
		return
	}
	defer writeFile.Close()

	_, err = writeFile.Write(file.Body)
	return

}

func (s *storage) SavePhoto(photo models.WritePhoto) (photoView models.Photo, err error) {
	fileName, err := s.createFileName(photo.Name)
	if err != nil {
		return
	}

	filePath, err := s.createFilePath(photo.Name)
	if err != nil {
		return
	}

	photoView.Name = fileName
	photoView.URL = s.photoPath + "/" + filePath + "/" + fileName

	writeFile, err := os.Create(photoView.URL)
	if err != nil {
		return
	}
	defer writeFile.Close()

	_, err = writeFile.Write(photo.Body)
	return
}

func (s *storage) InsertFile(file models.File) (fileID int, err error) {
	const sqlQuery = `
	INSERT INTO files (title, url)
	VALUES ($1, $2)
	RETURNING id;`

	err = s.db.QueryRow(sqlQuery, file.Name, file.URL).Scan(&fileID)
	return
}

func (s *storage) InsertPhoto(photo models.Photo) (photoID int, err error) {
	const sqlQuery = `
	INSERT INTO photo (title, url)
	VALUES ($1, $2)
	RETURNING id;`

	err = s.db.QueryRow(sqlQuery, photo.Name, photo.URL).Scan(&photoID)
	return
}

func (s *storage) createFilePath(name string) (pathName string, err error) {
	runes := []rune(name)
	if len(runes) < 3 {
		return pathName, errors.New("Некорректное имя пути")
	}

	return string(runes[:2]), nil
}

func (s *storage) createFileName(name string) (fileName string, err error) {
	h := md5.New()
	h.Write([]byte(time.Now().String() + name))
	h.Sum(nil)

	postfix, err := s.extractFormatFile(name)
	if err != nil {
		return
	}

	return string(h.Sum(nil)) + "." + postfix, nil
}

func (s *storage) extractFormatFile(fileName string) (postfix string, err error) {
	parts := strings.Split(fileName, ".")
	if len(parts) != 2 {
		return postfix, errors.New("Некорректное имя файла")
	}
	postfix = parts[1]
	return
}

//func (s *storage) SaveFiles(files []models.WriteFile) (fileViews []models.File, err error) {
//	for i, _ := range files {
//		fileView, err := s.SaveFile(files[i])
//		if err != nil {
//			return
//		}
//		fileViews = append(fileViews, fileView)
//	}
//
//	return
//}
