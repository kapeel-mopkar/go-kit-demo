package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/kapeel-mopkar/go-kit-demo/account"
	acctdb "github.com/kapeel-mopkar/go-kit-demo/account/db"
	"github.com/kapeel-mopkar/go-kit-demo/account/impl"

	_ "github.com/go-sql-driver/mysql"
)

const dbsource = "root:root@tcp(127.0.0.1:3306)/go-mysql"

func main() {

	var httpAddr = flag.String("http", ":8084", "http listen address")

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

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service stopped")

	var db *sql.DB
	{
		var err error
		db, err = sql.Open("mysql", dbsource)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	flag.Parse()

	ctx := context.Background()

	var svc account.Service
	{
		repository := acctdb.NewRepo(db, logger)

		svc = impl.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := account.MakeEndpoints(svc)

	go func() {
		fmt.Println("Listening on port ", *httpAddr)
		handler := account.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)

}
