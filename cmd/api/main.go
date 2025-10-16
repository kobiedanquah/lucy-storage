package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/primekobie/lucy/internal/handlers"
	"github.com/primekobie/lucy/internal/postgres"
	"github.com/primekobie/lucy/internal/services"
)

type ServerApplication struct {
	handler *handlers.ServiceHandler
}

func main() {

	_ = godotenv.Load()

	db, err := sql.Open("pgx", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	userStore := postgres.NewUserStore(db)
	userService := services.NewUserService(userStore)

	handler := handlers.NewServiceHandler(userService)

	app := &ServerApplication{
		handler: handler,
	}

	srv := http.Server{
		Addr:    ":8080",
		Handler: app.loadRoutes(),
	}

	log.Print("Starting Server")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
