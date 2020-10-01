package main

import (
	"database/sql"
	"github.com/BarniBl/SolAr_2020/cmd/handlers"
	postsHandler "github.com/BarniBl/SolAr_2020/cmd/handlers/posts"
	"github.com/BarniBl/SolAr_2020/internal/services/posts"
	fileStorage "github.com/BarniBl/SolAr_2020/internal/storages/uploadStorage"
	"github.com/BarniBl/SolAr_2020/internal/storages/postsStorage"
	"github.com/kelseyhightower/envconfig"
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

	db, err := sql.Open("postgres", cfg.DataBaseConnectionString)
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)

	fileStorage := fileStorage.NewStorage()

	postsStorage := postsStorage.NewStorage(db)
	postsService := posts.NewService(postsStorage)
	postsTransport := posts.NewTransport()

	postsHandler := postsHandler.NewHandler(postsService, postsTransport)

	middlewares := handlers.NewMiddleware()

	server := fasthttp.Server{
		Handler: handlers.NewFastHttpRouter(postsHandler, middlewares).Handler,
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
