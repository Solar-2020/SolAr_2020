package postStorage

import (
	"database/sql"
	"github.com/Solar-2020/SolAr_Backend_2020/internal/models"
	"strconv"
	"strings"
)

const (
	queryReturningID = "RETURNING id;"
)

type Storage interface {
	InsertPost(inputPost models.InputPost) (postID int, err error)

	UpdatePostStatus(postID int, status int) (err error)

	SelectFileIDs(postIDs []int) (matches []models.PostFileMatch, err error)
	SelectPhotoIDs(postIDs []int) (matches []models.PostPhotoMatch, err error)

	SelectPosts(request models.GetPostListRequest) (posts []models.InputPost, err error)
	SelectPayments(postIDs []int) (payments []models.Payment, err error)
	SelectInterviews(postIDs []int) (interviews []models.Interview, err error)
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) InsertPost(inputPost models.InputPost) (postID int, err error) {
	tx, err := s.db.Begin()
	if err != nil {
		return
	}
	defer tx.Rollback()

	postID, err = s.insertPost(tx, inputPost)
	if err != nil {
		return
	}

	err = s.insertPhotos(tx, inputPost.Photos, postID)
	if err != nil {
		return
	}

	err = s.insertFiles(tx, inputPost.Files, postID)
	if err != nil {
		return
	}

	err = tx.Commit()

	return
}

func (s *storage) insertPost(tx *sql.Tx, post models.InputPost) (postID int, err error) {
	const sqlQuery = `
	INSERT INTO posts(create_by, publish_date, group_id, text)
	VALUES ($1, $2, $3, $4)
	RETURNING id;`

	err = tx.QueryRow(sqlQuery, post.CreateBy, post.PublishDate, post.GroupID, post.Text).Scan(&postID)

	return
}

func (s *storage) insertPhotos(tx *sql.Tx, photos []int, postID int) (err error) {
	if len(photos) == 0 {
		return
	}

	const sqlQueryTemplate = `
	INSERT INTO photos(post_id, photo_id)
	VALUES `

	var params []interface{}

	sqlQuery := sqlQueryTemplate + s.createInsertQuery(len(photos), 2)

	for i, _ := range photos {
		params = append(params, postID, photos[i])
	}

	for i := 1; i <= len(photos)*2; i++ {
		sqlQuery = strings.Replace(sqlQuery, "?", "$"+strconv.Itoa(i), 1)
	}

	_, err = tx.Exec(sqlQuery, params...)

	return
}

func (s *storage) insertFiles(tx *sql.Tx, files []int, postID int) (err error) {
	if len(files) == 0 {
		return
	}

	const sqlQueryTemplate = `
	INSERT INTO files(post_id, file_id)
	VALUES `

	var params []interface{}

	sqlQuery := sqlQueryTemplate + s.createInsertQuery(len(files), 2)

	for i, _ := range files {
		params = append(params, postID, files[i])
	}

	for i := 1; i <= len(files)*2; i++ {
		sqlQuery = strings.Replace(sqlQuery, "?", "$"+strconv.Itoa(i), 1)
	}

	_, err = tx.Exec(sqlQuery, params...)

	return
}

func (s *storage) UpdatePostStatus(postID int, status int) (err error) {
	const sqlQuery = `
	UPDATE posts
	SET status_id = $2
	WHERE id = $1;`

	_, err = s.db.Exec(sqlQuery, postID, status)

	return
}

func (s *storage) SelectPosts(request models.GetPostListRequest) (posts []models.InputPost, err error) {
	sqlQuery := `
	SELECT p.id, p.text, p.group_id, p.publish_date
	FROM posts.posts AS p
	WHERE p.create_by = $1
	  AND p.group_id = $2
	  AND p.status_id = 2
	  AND p.publish_date <= $3
	ORDER BY p.publish_date DESC
	LIMIT $4`

	rows, err := s.db.Query(sqlQuery, request.UserID, request.GroupID, request.StartFrom, request.Limit)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var tempPost models.InputPost
		err = rows.Scan(&tempPost.ID, &tempPost.Text, &tempPost.GroupID, &tempPost.PublishDate)
		if err != nil {
			return
		}
		posts = append(posts, tempPost)
	}

	return
}

func (s *storage) SelectInterviews(postIDs []int) (interviews []models.Interview, err error) {
	const sqlQueryTemplate = `
	SELECT i.id, i.text, i.type, i.post_id
	FROM interview AS i
	WHERE i.post_id IN `

	sqlQuery := sqlQueryTemplate + createIN(len(postIDs))

	var params []interface{}

	for i, _ := range postIDs {
		params = append(params, postIDs[i])
	}

	for i := 1; i <= len(postIDs)*1; i++ {
		sqlQuery = strings.Replace(sqlQuery, "?", "$"+strconv.Itoa(i), 1)
	}

	rows, err := s.db.Query(sqlQuery, params...)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var tempInterview models.Interview
		err = rows.Scan(&tempInterview.ID, &tempInterview.Text, &tempInterview.Type, &tempInterview.PostID)
		if err != nil {
			return
		}
		interviews = append(interviews, tempInterview)
	}

	// TODO SELECT ANSWERS

	return
}

func (s *storage) SelectPayments(postIDs []int) (payments []models.Payment, err error) {
	const sqlQueryTemplate = `
	SELECT p.id, p.cost, p.currency_id, p.post_id
	FROM payments AS p
	WHERE i.post_id IN `

	sqlQuery := sqlQueryTemplate + createIN(len(postIDs))

	var params []interface{}

	for i, _ := range postIDs {
		params = append(params, postIDs[i])
	}

	for i := 1; i <= len(postIDs)*1; i++ {
		sqlQuery = strings.Replace(sqlQuery, "?", "$"+strconv.Itoa(i), 1)
	}

	rows, err := s.db.Query(sqlQuery, params...)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var tempPayment models.Payment
		err = rows.Scan(&tempPayment.ID, &tempPayment.Cost, &tempPayment.Currency, &tempPayment.PostID)
		if err != nil {
			return
		}
		payments = append(payments, tempPayment)
	}

	return
}

func (s *storage) SelectFileIDs(postIDs []int) (matches []models.PostFileMatch, err error) {
	const sqlQueryTemplate = `
	SELECT f.post_id, f.file_id
	FROM files AS f
	WHERE f.post_id IN `

	sqlQuery := sqlQueryTemplate + createIN(len(postIDs))

	var params []interface{}

	for i, _ := range postIDs {
		params = append(params, postIDs[i])
	}

	for i := 1; i <= len(postIDs)*1; i++ {
		sqlQuery = strings.Replace(sqlQuery, "?", "$"+strconv.Itoa(i), 1)
	}

	rows, err := s.db.Query(sqlQuery, params...)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var tempMatch models.PostFileMatch
		err = rows.Scan(&tempMatch.PostID, &tempMatch.FileID)
		if err != nil {
			return
		}
		matches = append(matches, tempMatch)
	}

	return
}

func (s *storage) SelectPhotoIDs(postIDs []int) (matches []models.PostPhotoMatch, err error) {
	const sqlQueryTemplate = `
	SELECT p.post_id, p.photo_id
	FROM photos AS p
	WHERE p.post_id IN `

	sqlQuery := sqlQueryTemplate + createIN(len(postIDs))

	var params []interface{}

	for i, _ := range postIDs {
		params = append(params, postIDs[i])
	}

	for i := 1; i <= len(postIDs)*1; i++ {
		sqlQuery = strings.Replace(sqlQuery, "?", "$"+strconv.Itoa(i), 1)
	}

	rows, err := s.db.Query(sqlQuery, params...)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var tempMatch models.PostPhotoMatch
		err = rows.Scan(&tempMatch.PostID, &tempMatch.PhotoID)
		if err != nil {
			return
		}
		matches = append(matches, tempMatch)
	}

	return
}

func createIN(count int) (queryIN string) {
	queryIN = "("
	for i := 0; i < count; i++ {
		queryIN += "?, "
	}
	queryINRune := []rune(queryIN)
	queryIN = string(queryINRune[:len(queryINRune)-2])
	queryIN += ")"
	return
}

func (s *storage) createInsertQuery(sliceLen int, structLen int) (query string) {
	query = ""
	for i := 0; i < sliceLen; i++ {
		query += "("
		for j := 0; j < structLen; j++ {
			query += "?,"
		}
		// delete last comma
		query = strings.TrimRight(query, ",")
		query += "),"
	}
	// delete last comma
	query = strings.TrimRight(query, ",")

	return
}
