package main

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/kobiedanquah/lucy/internal/handlers"
	"github.com/kobiedanquah/lucy/internal/mailer"
	"github.com/kobiedanquah/lucy/internal/postgres"
	"github.com/kobiedanquah/lucy/internal/services"
)

type ServerApplication struct {
	handler *handlers.ServiceHandler
}

type Config struct {
	MailConfig    *mailer.Config
	PostgresURL   string
	ServerAddress string
}

func loadConfig() *Config {

	mailCfg := &mailer.Config{
		Host:        os.Getenv("MAIL_HOST"),
		Token:       os.Getenv("MAIL_TOKEN"),
		SenderEmail: os.Getenv("SENDER_EMAIL"),
		SenderName:  os.Getenv("SENDER_NAME"),
	}

	return &Config{
		MailConfig:    mailCfg,
		PostgresURL:   os.Getenv("DB_URL"),
		ServerAddress: os.Getenv("PORT"),
	}
}

func main() {

	_ = godotenv.Load()

	setupLogging()

	config := loadConfig()

	db, err := sql.Open("pgx", config.PostgresURL)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	mailer := mailer.NewMailer(config.MailConfig)
	userService := services.NewUserService(postgres.NewUserStore(db), mailer)

	handler := handlers.NewServiceHandler(userService)

	app := &ServerApplication{
		handler: handler,
	}

	srv := http.Server{
		Addr:    ":8080",
		Handler: app.loadRoutes(),
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	serverErr := make(chan error, 1)
	go func() {
		slog.Info("Starting Server")
		serverErr <- srv.ListenAndServe()
	}()

	select {
	case err := <-serverErr:
		if err != nil {
			panic(err)
		}
	case sig := <-stop:
		slog.Info("Shutting down server", "signal", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			slog.Error("Graceful shutdown failed", "error", err)
		}
	}
}
