package posts

import (
	"database/sql"
	"fmt"
	"github.com/Solar-2020/GoUtils/context"
	"github.com/Solar-2020/SolAr_Backend_2020/internal/models"
	interviewSt "github.com/Solar-2020/SolAr_Backend_2020/internal/storages/interviewStorage"
	paymentSt "github.com/Solar-2020/SolAr_Backend_2020/internal/storages/paymentStorage"
	postSt "github.com/Solar-2020/SolAr_Backend_2020/internal/storages/postStorage"
	uploadSt "github.com/Solar-2020/SolAr_Backend_2020/internal/storages/uploadStorage"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
	"time"
)

func TestPosts(t *testing.T) {
	postsDB, err := sql.Open("postgres", os.Getenv("POSTS_DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatal(err)
		return
	}

	postsDB.SetMaxIdleConns(5)
	postsDB.SetMaxOpenConns(10)

	uploadDB, err := sql.Open("postgres", os.Getenv("UPLOAD_DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatal(err)
	}

	postsDB.SetMaxIdleConns(5)
	postsDB.SetMaxOpenConns(10)

	uploadStorage := uploadSt.NewStorage("", "", uploadDB)

	interviewStorage := interviewSt.NewStorageProxy(os.Getenv("INTERVIEW_SERVICE_ADDRESS"))
	paymentStorage := paymentSt.NewStorage(postsDB)

	postStorage := postSt.NewStorage(postsDB)
	postsService := NewService(postStorage, uploadStorage, interviewStorage, paymentStorage, os.Getenv("GROUP_SERVICE_ADDRESS"))
	testData := models.GetPostListRequest{
		UserID:    12,
		GroupID:   35,
		Limit:     10,
		StartFrom: time.Now(),
	}

	ctx := context.Context{
		RequestCtx: nil,
		Session:    nil,
	}

	list, err := postsService.GetList(ctx, testData)
	fmt.Println(list)
}
