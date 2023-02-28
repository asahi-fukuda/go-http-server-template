package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/myuoncorp/go-http-server-template/configuration"
	"github.com/myuoncorp/go-http-server-template/infra/repository"
	"github.com/myuoncorp/go-http-server-template/logger"
	"github.com/myuoncorp/go-http-server-template/oapistub"
	"github.com/myuoncorp/go-http-server-template/server"
	"github.com/myuoncorp/go-http-server-template/usecase/organization"

	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"

	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/logrusadapter"
	"github.com/sirupsen/logrus"
)

var (
	version   = "unknown"
	revision  = "unknown"
	builtAt   = "unknown"
	goVersion = "unknown"
)

func main() {
	logger.Infof("Version: %s", version)
	logger.Infof("Revision: %s", revision)
	logger.Infof("Built: %s", builtAt)
	logger.Infof("Go: %s", goVersion)
	configuration.Load()
	conf := configuration.Get()

	// --------------------------------------------------
	// Connect to database
	// --------------------------------------------------
	dsn := conf.DSN
	if dsn == "" {
		logger.Fatal("DSN is not specified")
	}
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		logger.Fatal(err)
	}

	// --------------------------------------------------
	// Create repository and service
	// --------------------------------------------------
	repo := repository.NewOrganizationRepository(db)

	// --------------------------------------------------
	// Setup the application
	// --------------------------------------------------
	s := &server.APIServer{}
	s.Controllers(
		&server.HealthCheckController{
			Version:   version,
			Revision:  revision,
			BuiltAt:   builtAt,
			GoVersion: goVersion,
		},
		&server.OrganizationsController{
			ListOrganizationsUseCase: &organization.List{
				OrganizationRepository: repo,
			},
			GetOrganizationsUseCase: &organization.Get{
				OrganizationRepository: repo,
			},
		},
	)

	// --------------------------------------------------
	// Setup echo server
	// --------------------------------------------------
	swagger, err := oapistub.GetSwagger()
	if err != nil {
		panic(err)
	}
	swagger.Servers = nil

	e := echo.New()

	// --------------------------------------------------
	// Init sentry
	// --------------------------------------------------
	err = initSentry(conf.SentryDSN, conf.Env, releaseName(), conf.Env != "production") // TODO: 本番では失敗した際のハンドリングをした方が良い
	logger.Infof("sentry error=%v", err)

	e.Use(sentryecho.New(sentryecho.Options{
		Repanic: true,
	}))
	defer sentry.Flush(3 * time.Second)

	e.Use(echomiddleware.LoggerWithConfig(echomiddleware.LoggerConfig{
		Skipper: func(ctx echo.Context) bool {
			return ctx.Path() == "/"
		},
		Format:           echomiddleware.DefaultLoggerConfig.Format,
		CustomTimeFormat: echomiddleware.DefaultLoggerConfig.CustomTimeFormat,
	}))
	e.Use(echomiddleware.CORS())
	e.Use(middleware.OapiRequestValidator(swagger))

	oapistub.RegisterHandlers(e, s)

	// --------------------------------------------------
	// Start the server in the background
	// --------------------------------------------------
	port := configuration.Get().Port
	go func() {
		e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%s", port)))
	}()

	// --------------------------------------------------
	// Wait signal for graceful stop
	// --------------------------------------------------
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("stopping server...")
	e.Shutdown(context.Background())
	logger.Info("server exiting")
}

func openDB(env, driver, dsn string) *sqlx.DB {
	inner, err := sql.Open(driver, dsn)
	if err != nil {
		logger.Fatal(err)
	}
	if env == "local" { // ENV が local であれば実行される SQL ログを出す。
		sqlLogger := logrus.New()
		sqlLogger.SetOutput(os.Stdout)
		sqlLogger.SetLevel(logrus.DebugLevel)
		sqlLogger.SetFormatter(&logrus.TextFormatter{
			DisableQuote: true,
		})
		db := sqldblogger.OpenDriver(
			dsn,
			inner.Driver(),
			logrusadapter.New(sqlLogger),
		)
		return sqlx.NewDb(db, dsn)
	}
	return sqlx.NewDb(inner, dsn)
}

func releaseName() string {
	return fmt.Sprintf("%s-%s-%s-%s", version, revision, builtAt, goVersion)
}

func initSentry(dsn, environment, release string, debug bool) error {
	if dsn == "" || environment == "" || release == "" {
		return fmt.Errorf("sentry option empty. dsn=%q, environment=%q, release=%q, debug=%t", dsn, environment, release, debug)
	}

	return sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		Environment:      environment,
		Release:          release,
		TracesSampleRate: 1.0,
		AttachStacktrace: true,
		Debug:            debug,
	})
}
