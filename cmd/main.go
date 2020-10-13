package main

import (
	"database/sql"
	"github.com/Solar-2020/SolAr_2020/cmd/handlers"
	postsHandler "github.com/Solar-2020/SolAr_2020/cmd/handlers/posts"
	uploadHandler "github.com/Solar-2020/SolAr_2020/cmd/handlers/upload"
	"github.com/Solar-2020/SolAr_2020/internal/errorWorker"
	"github.com/Solar-2020/SolAr_2020/internal/services/posts"
	"github.com/Solar-2020/SolAr_2020/internal/services/upload"
	"github.com/Solar-2020/SolAr_2020/internal/storages/postsStorage"
	"github.com/Solar-2020/SolAr_2020/internal/storages/uploadStorage"
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

	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}

	postsDB, err := sql.Open("postgres", cfg.PostsDataBaseConnectionString)
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}

	postsDB.SetMaxIdleConns(5)
	postsDB.SetMaxOpenConns(10)

	uploadDB, err := sql.Open("postgres", cfg.UploadDataBaseConnectionString)
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}

	postsDB.SetMaxIdleConns(5)
	postsDB.SetMaxOpenConns(10)

	//userDB, err := sql.Open("postgres", cfg.UserDataBaseConnectionString)
	//if err != nil {
	//	log.Fatal().Msg(err.Error())
	//	return
	//}

	//userDB.SetMaxIdleConns(5)
	//userDB.SetMaxOpenConns(10)

	errorWorker := errorWorker.NewErrorWorker()

	uploadStorage := uploadStorage.NewStorage(cfg.PhotoPath, cfg.FilePath, uploadDB)
	uploadService := upload.NewService(uploadStorage)
	uploadTransport := upload.NewTransport()
	uploadHandler := uploadHandler.NewHandler(uploadService, uploadTransport, errorWorker)

	postsStorage := postsStorage.NewStorage(postsDB)
	postsService := posts.NewService(postsStorage, uploadStorage)
	postsTransport := posts.NewTransport()

	postsHandler := postsHandler.NewHandler(postsService, postsTransport, errorWorker)

	middlewares := handlers.NewMiddleware()

	server := fasthttp.Server{
		Handler: handlers.NewFastHttpRouter(postsHandler, uploadHandler, middlewares).Handler,
	}

	go func() {
		log.Info().Str("msg", "start server").Str("port", cfg.Port).Send()
		if err := server.ListenAndServe(":" + cfg.Port); err != nil {
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
