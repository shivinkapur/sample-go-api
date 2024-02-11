package persistence

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"strings"
	"sync"

	_ "github.com/lib/pq"

	"github.com/shivinkapur/sample-go-api/persistence/entities"
)

var repoInitOnce sync.Once
var REPOSITORY Repository

func GetRepository() Repository {

	repoInitOnce.Do(func() {
		const dsnKey = "DB_CONNECTIONSTRING"
		dsn := os.Getenv(dsnKey)

		if strings.TrimSpace(dsn) == "" {
			log.Fatalf("FATAL: environment variable %s not set", dsnKey)
		}

		db, err := sql.Open("postgres", dsn)

		if err != nil {
			log.Fatalf("%+v", err)
		}

		REPOSITORY, err = NewDBRepository(db)

		if err != nil {
			log.Fatalf("%+v", err)
		}
	})

	return REPOSITORY
}

type Repository interface {
	GetUserByUserName(username string) (*entities.User, error)
}

func NewDBRepository(db *sql.DB) (Repository, error) {
	return &dbRepository{db: db}, nil
}

type dbRepository struct {
	Repository
	db *sql.DB
}

var (
	ErrUserNotFound = errors.New("user was not found")
)

func (r dbRepository) GetUserByUserName(username string) (*entities.User, error) {
	var result entities.User

	if username == "shivin" {
		return nil, ErrUserNotFound
	}

	result = entities.User{
		Id:        "1",
		FirstName: "Jon",
		LastName:  "Doe",
	}

	return &result, nil
}
