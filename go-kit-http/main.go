package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/muhammadarash1997/go-kit-http/endpoints"
	"github.com/muhammadarash1997/go-kit-http/repositories"
	"github.com/muhammadarash1997/go-kit-http/services"
	"github.com/muhammadarash1997/go-kit-http/transports"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/log/level"

	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const dbsource = "postgresql://postgres:postgres@localhost:5432/gokitexample?sslmode=disable"

func main() {
	var httpAddress = flag.String("http", ":8080", "HTTP Listen Address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("message", "Service Started")
	defer level.Info(logger).Log("message", "Service Ended")

	var db *sql.DB
	var err error
	{
		db, err = sql.Open("postgres", dbsource)
		if err != nil {
			level.Error(logger).Log("Exit", err)
			os.Exit(-1)
		}
	}

	flag.Parse()
	ctx := context.Background()

	var repository repositories.Repository
	var service services.Service
	{
		repository = repositories.NewRepository(db, logger)
		service = services.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := endpoints.MakeEndpoints(service)

	go func() {
		fmt.Println("Listening on Port", *httpAddress)
		handler := transports.NewHTTPServer(ctx, endpoints)

		// Listen and Serve
		errs <- http.ListenAndServe(*httpAddress, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
