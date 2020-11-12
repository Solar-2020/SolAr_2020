package main

import (
	"database/sql"
	asapi "github.com/Solar-2020/Account-Backend/pkg/api"
	authapi "github.com/Solar-2020/Authorization-Backend/pkg/api"
	"github.com/Solar-2020/GoUtils/context/session"
	"github.com/Solar-2020/GoUtils/http/errorWorker"
	"github.com/Solar-2020/SolAr_Backend_2020/cmd/config"
	"github.com/Solar-2020/SolAr_Backend_2020/cmd/handlers"
	postsHandler "github.com/Solar-2020/SolAr_Backend_2020/cmd/handlers/posts"
	uploadHandler "github.com/Solar-2020/SolAr_Backend_2020/cmd/handlers/upload"
	"github.com/Solar-2020/SolAr_Backend_2020/internal/clients/account"
	"github.com/Solar-2020/SolAr_Backend_2020/internal/clients/auth"
	"github.com/Solar-2020/SolAr_Backend_2020/internal/clients/group"
	"github.com/Solar-2020/SolAr_Backend_2020/internal/clients/interview"
	"github.com/Solar-2020/SolAr_Backend_2020/internal/services/posts"
	"github.com/Solar-2020/SolAr_Backend_2020/internal/services/upload"
	"github.com/Solar-2020/SolAr_Backend_2020/internal/storages/paymentStorage"
	"github.com/Solar-2020/SolAr_Backend_2020/internal/storages/postStorage"
	"github.com/Solar-2020/SolAr_Backend_2020/internal/storages/uploadStorage"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})

	err := envconfig.Process("", &config.Config)
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}

	postsDB, err := sql.Open("postgres", config.Config.PostsDataBaseConnectionString)
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}

	postsDB.SetMaxIdleConns(5)
	postsDB.SetMaxOpenConns(10)

	uploadDB, err := sql.Open("postgres", config.Config.UploadDataBaseConnectionString)
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}

	uploadDB.SetMaxIdleConns(5)
	uploadDB.SetMaxOpenConns(10)

	//userDB, err := sql.Open("postgres", cfg.UserDataBaseConnectionString)
	//if err != nil {
	//	log.Fatal().Msg(err.Error())
	//	return
	//}

	//userDB.SetMaxIdleConns(5)
	//userDB.SetMaxOpenConns(10)
	groupClient := group.NewClient(config.Config.GroupServiceAddress, config.Config.ServerSecret)

	errorWorker := errorWorker.NewErrorWorker()

	uploadStorage := uploadStorage.NewStorage(config.Config.PhotoPath, config.Config.FilePath, uploadDB)
	uploadService := upload.NewService(uploadStorage)
	uploadTransport := upload.NewTransport()
	uploadHandler := uploadHandler.NewHandler(uploadService, uploadTransport, errorWorker)

	//interviewStorage := interviewStorage.NewStorage(postsDB)
	accountClient := account.NewClient(config.Config.AccountServiceAddress, config.Config.ServerSecret)

	interviewStorage := interview.NewClient(config.Config.InterviewService, config.Config.ServerSecret)
	paymentStorage := paymentStorage.NewStorage(postsDB)
	postStorage := postStorage.NewStorage(postsDB)
	postsService := posts.NewService(postStorage, uploadStorage, interviewStorage, paymentStorage, groupClient, accountClient)
	postsTransport := posts.NewTransport()

	postsHandler := postsHandler.NewHandler(postsService, postsTransport, errorWorker)

	authClient := auth.NewClient(config.Config.AuthServiceAddress, config.Config.ServerSecret)

	middlewares := handlers.NewMiddleware(&log, authClient)

	initServices()

	server := fasthttp.Server{
		Handler: handlers.NewFastHttpRouter(postsHandler, uploadHandler, middlewares).Handler,
	}

	go func() {
		log.Info().Str("msg", "start server").Str("port", config.Config.Port).Send()
		if err := server.ListenAndServe(":" + config.Config.Port); err != nil {
			log.Error().Str("msg", "server run failure").Err(err).Send()
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	defer func(sig os.Signal) {

		log.Info().Str("msg", "received signal, exiting").Str("signal", sig.String()).Send()

		if err := server.Shutdown(); err != nil {
			log.Error().Str("msg", "server shutdown failure").Err(err).Send()
		}

		//dbConnection.Shutdown()
		log.Info().Str("msg", "goodbye").Send()
	}(<-c)
}

func initServices() {
	authService := authapi.AuthClient{
		Addr: config.Config.AuthServiceAddress,
	}
	session.RegisterAuthService(&authService)
	accountService := asapi.AccountClient{
		Addr: config.Config.AccountServiceAddress,
	}
	session.RegisterAccountService(&accountService)
}
