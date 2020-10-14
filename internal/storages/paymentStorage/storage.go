package paymentStorage

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
	InsertPayments(payments []models.Payment, postID int) (err error)
	SelectPayments(postIDs []int) (payments []models.Payment, err error)
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) InsertPayments(payments []models.Payment, postID int) (err error) {
	if len(payments) == 0 {
		return
	}

	const sqlQueryTemplate = `
	INSERT INTO payments(post_id, cost, currency_id)
	VALUES `

	var params []interface{}

	sqlQuery := sqlQueryTemplate + s.createInsertQuery(len(payments), 3) + queryReturningID

	for i, _ := range payments {
		params = append(params, postID, payments[i].Cost, payments[i].Currency)
	}

	for i := 1; i <= len(payments)*3; i++ {
		sqlQuery = strings.Replace(sqlQuery, "?", "$"+strconv.Itoa(i), 1)
	}

	_, err = s.db.Exec(sqlQuery, params...)

	return
}

func (s *storage) SelectPayments(postIDs []int) (payments []models.Payment, err error) {
	const sqlQueryTemplate = `
	SELECT p.id, p.cost, p.currency_id, p.post_id
	FROM payments AS p
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
		var tempPayment models.Payment
		err = rows.Scan(&tempPayment.ID, &tempPayment.Cost, &tempPayment.Currency, &tempPayment.PostID)
		if err != nil {
			return
		}
		payments = append(payments, tempPayment)
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
		query =  strings.TrimRight(query, ",")
		query += "),"
	}
	// delete last comma
	query =  strings.TrimRight(query, ",")

	return
}