package main

import (
	"database/sql"
	"log"
	"log/slog"
	"net/http"
	"os"

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

	slog.Info("Starting Server")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
