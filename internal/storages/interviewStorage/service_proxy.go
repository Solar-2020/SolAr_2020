package interviewStorage

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Solar-2020/SolAr_Backend_2020/internal/models"
	//"github.com/valyala/fasthttp"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	//"time"
)

type ServiceStorage struct {
	serviceAddress string
}

func NewStorageProxy(serviceAddress string) Storage {
	return &ServiceStorage{
		serviceAddress: serviceAddress,
	}
}

//func (s *ServiceStorage) InsertInterviews(interviews []models.Interview, postID int) (err error) {
//	if len(interviews) == 0 {
//		return
//	}
//
//	const sqlQuery = `
//	INSERT INTO interviews(text, type, post_id)
//	VALUES ($1, $2, $3)
//	RETURNING id;`
//
//	tx, err := s.db.Begin()
//	if err != nil {
//		return
//	}
//	defer tx.Rollback()
//
//	for i, _ := range interviews {
//		var currentInterviewID int
//		err = s.db.QueryRow(sqlQuery, interviews[i].Text, interviews[i].Type, postID).Scan(&currentInterviewID)
//		if err != nil {
//			return
//		}
//
//		err = s.insertAnswers(tx, interviews[i].Answers, currentInterviewID)
//		if err != nil {
//			return
//		}
//	}
//
//	err = tx.Commit()
//
//	return
//}

func (s *ServiceStorage) InsertInterviews(interviews []models.Interview, postID int) (err error) {
	request := struct{
		Interviews []models.Interview `json:"interviews"`
		PostID int `json:"postID"`
	}{
		interviews,
		postID,
	}
	body, err := json.Marshal(request)
	if err != nil {
		return
	}

	resp, err := http.Post(s.serviceAddress+"/interview/create", "application/json", bytes.NewReader(body))
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad response: %d", resp.StatusCode)
		return
	}
	//var resInterview []models.Interview
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(body, &request)
	interviews = request.Interviews
	return err

}

func (s *ServiceStorage) insertAnswers(tx *sql.Tx, answers []models.Answer, interviewID int) (err error) {
	if len(answers) == 0 {
		return
	}

	sqlQueryTemplate := `
	INSERT INTO answers(interview_id, text)
	VALUES `

	if len(answers) == 0 {
		return
	}

	var params []interface{}

	sqlQuery := sqlQueryTemplate + s.createInsertQuery(len(answers), 2)

	for i, _ := range answers {
		params = append(params, interviewID, answers[i].Text)
	}

	for i := 1; i <= len(answers)*2; i++ {
		sqlQuery = strings.Replace(sqlQuery, "?", "$"+strconv.Itoa(i), 1)
	}

	_, err = tx.Exec(sqlQuery, params...)
	return
}

func (s *ServiceStorage) SelectInterviews(postIDs []int) (interviews []models.Interview, err error) {
	body, err := json.Marshal(postIDs)
	if err != nil {
		return
	}

	resp, err := http.Post(s.serviceAddress+"/interview", "application/json", bytes.NewReader(body))
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad response: %d", resp.StatusCode)
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(body, &interviews)
	return
}

func (s *ServiceStorage) selectAnswers(interviewIDs []int) (answers []models.Answer, err error) {
	//const sqlQueryTemplate = `
	//SELECT a.id, a.text, a.interview_id
	//FROM answers AS a
	//WHERE a.interview_id IN `
	//
	//sqlQuery := sqlQueryTemplate + createIN(len(interviewIDs))
	//
	//var params []interface{}
	//
	//for i, _ := range interviewIDs {
	//	params = append(params, interviewIDs[i])
	//}
	//
	//for i := 1; i <= len(interviewIDs)*1; i++ {
	//	sqlQuery = strings.Replace(sqlQuery, "?", "$"+strconv.Itoa(i), 1)
	//}
	//
	//rows, err := s.db.Query(sqlQuery, params...)
	//if err != nil {
	//	return
	//}
	//defer rows.Close()
	//
	//for rows.Next() {
	//	var tempAnswer models.Answer
	//	err = rows.Scan(&tempAnswer.ID, &tempAnswer.Text, &tempAnswer.InterviewID)
	//	if err != nil {
	//		return
	//	}
	//	answers = append(answers, tempAnswer)
	//}

	return
}

func (s *ServiceStorage) createInsertQuery(sliceLen int, structLen int) (query string) {
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

