package persistence

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
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
	AddUser(user *entities.User) (*entities.User, error)
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
	ErrUserNotFound      = errors.New("user was not found")
	ErrDBError           = errors.New("database error")
	ErrUserAlreadyExists = errors.New("user already exists")
)

func (r dbRepository) GetUserByUserName(username string) (*entities.User, error) {
	var result entities.User

	if username == "shivin" {
		return nil, ErrUserNotFound
	}

	q := `SELECT id, first_name, last_name, username, email, password, phone, user_status, deleted, created_at, modified_at FROM PET_STORE.USERS WHERE username = $1`

	err := r.db.QueryRow(q, username).Scan(&result.Id, &result.FirstName, &result.LastName, &result.Username, &result.Email, &result.Password, &result.Phone, &result.UserStatus, &result.Deleted, &result.CreatedAt, &result.ModifiedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, ErrDBError
	}

	return &result, nil
}

func (r dbRepository) AddUser(user *entities.User) (*entities.User, error) {

	// check if user already exists
	_, err := r.GetUserByUserName(user.Username)

	if err == nil {
		return nil, ErrUserAlreadyExists
	}

	user.CreatedAt = time.Now().UTC().Unix()
	user.ModifiedAt = user.CreatedAt
	user.Deleted = false
	user.Id = uuid.NewString()

	// set user status to active
	user.UserStatus = 1

	// if no password, set password to default value
	if user.Password == "" {
		user.Password = "password"
	}

	q := `INSERT INTO PET_STORE.USERS (id, first_name, last_name, username, email, password, phone, user_status, deleted, created_at, modified_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err = r.db.Exec(q, user.Id, user.FirstName, user.LastName, user.Username, user.Email, user.Password, user.Phone, user.UserStatus, user.Deleted, user.CreatedAt, user.ModifiedAt)

	if err != nil {
		return nil, ErrDBError
	}

	return user, nil
}
