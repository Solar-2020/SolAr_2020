package postsStorage

import (
	"database/sql"
	"github.com/BarniBl/SolAr_2020/internal/models"
)

type Storage interface {
	InsertPost(models.InputPost) (post models.InputPost, err error)
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) InsertPost(models.InputPost) (post models.InputPost, err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	post.ID, err = s.insertPost(post)
	if err != nil {
		return
	}

	s.insertInterviews(post.Interviews)

	s.insertPayments(post.Payments)
}

func (s *storage) insertPost(post models.InputPost) (postID int, err error) {
	query := `
	INSERT INTO
	`
}

func (s *storage) insertInterviews(interviews []models.Interview) (err error) {
	query := `
	INSERT INTO
	`
}

func (s *storage) insertPayments(payments []models.Payment) (err error) {
	query := `
	INSERT INTO
	`
}
