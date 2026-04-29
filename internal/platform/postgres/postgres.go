package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	// "github.com/google/uuid"
	// "github.com/kungrem23/quizgo/internal/store/models"
	// "github.com/kungrem23/quizgo/internal/store/repos"
	_ "github.com/lib/pq"
)

const (
	pgHost = "localhost"
	port   = "5432"
	user   = "danilmitrosin"
	dbname = "quizgo"
)

func ConnectPG() *sql.DB {
	pgCfg := fmt.Sprintf("host=%s port=%s user=%s dbname=%s"+
		" sslmode=disable", pgHost, port, user, dbname)
	// log.Println("Connecting to PG...")
	db, err := sql.Open("postgres", pgCfg)
	if err != nil {
		log.Fatalf("Connecting to PG error: %v", err)
	}
	log.Println("Connected to PG succesfully")
	return db
}

func ApplyMigrations(db *sql.DB) {
	files, err := filepath.Glob("../../internal/platform/postgres/migrations/*.sql")
	if err != nil {
		log.Fatalf("Migrations file error: %v", err)
	}
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("Reading file error: %v", err)
		}
		_, err = db.Exec(string(content))
		if err != nil {
			log.Fatalf("Applying migration error: %v", err)
		}
		log.Printf("Migration %v applied succesfully\n", file)
	}
	log.Println("All migrations applied succesfully")
}

func NewDBConnection() *sql.DB {
	db := ConnectPG()
	ApplyMigrations(db)
	return db
}
