package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

  "github.com/Tecu23/snipperbox/internal/models"

  _ "github.com/go-sql-driver/mysql"
)

type application struct {
  logger *slog.Logger
  snippets *models.SnippetModel
}

func main() {
  addr :=  flag.String("addr", ":4000", "HTTP network address")

  // Define a new command-line flag for the MySQL DSN string. 
  dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
  flag.Parse()

  logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

  // Database connection 
  db, err := openDB(*dsn)
  if err != nil {
    logger.Error(err.Error())
    os.Exit(1)
  }

  // We also defer a call to db.Close(), so that the connection pool is closed
  // before the main() function exists
  defer db.Close()

  // Initialize a new instance of out application struct
  app := &application{
    logger: logger,
    snippets: &models.SnippetModel{DB: db},
  }

  logger.Info("Starting server", "addr", *addr)

  err = http.ListenAndServe(*addr, app.routes())

  logger.Error(err.Error())
  os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
  db, err := sql.Open("mysql", dsn)
  if err != nil {
    return nil, err
  }

  err = db.Ping()
  if err != nil {
    db.Close()
    return nil, err
  }

  return db, nil
}
