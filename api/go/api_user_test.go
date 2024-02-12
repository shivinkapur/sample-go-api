package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shivinkapur/sample-go-api/persistence"
	"github.com/shivinkapur/sample-go-api/persistence/entities"
)

func TestUserAPI_GetUserByNameShivin(t *testing.T) {
	// Create a new Gin context for testing
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	t.Setenv("DB_CONNECTIONSTRING", "something")

	// Set up the test case
	username := "shivin"
	c.AddParam(PATH_PARAM_USERNAME, username)

	// Create a new instance of UserAPI
	api := &UserAPI{}

	// Call the GetUserByName function
	api.GetUserByName(c)

	// Check the response status code
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, w.Code)
	}
}

func TestUserAPI_GetUserByNameUserNotFoundError(t *testing.T) {
	// Create a new Gin context for testing
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	t.Setenv("DB_CONNECTIONSTRING", "something")

	// Set up the test case
	username := "error"
	c.AddParam(PATH_PARAM_USERNAME, username)

	// Mock the repository
	oldRepository := persistence.REPOSITORY
	persistence.REPOSITORY = persistence.GetRepository()
	persistence.REPOSITORY = MockRepository{
		Repository: oldRepository,
		err:        persistence.ErrUserNotFound,
	}

	// Create a new instance of UserAPI
	api := &UserAPI{}

	// Call the GetUserByName function
	api.GetUserByName(c)
	persistence.REPOSITORY = oldRepository // Reset the repository

	// Check the response status code
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, w.Code)
	}
}

func TestUserAPI_GetUserByNameInternalServerError(t *testing.T) {
	// Create a new Gin context for testing
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	t.Setenv("DB_CONNECTIONSTRING", "something")

	// Set up the test case
	username := "whatever"
	c.AddParam(PATH_PARAM_USERNAME, username)

	// Mock the repository
	oldRepository := persistence.REPOSITORY
	persistence.REPOSITORY = persistence.GetRepository()
	persistence.REPOSITORY = MockRepository{
		Repository: oldRepository,
		err:        errors.New("some error"),
	}

	// Create a new instance of UserAPI
	api := &UserAPI{}

	// Call the GetUserByName function
	api.GetUserByName(c)
	persistence.REPOSITORY = oldRepository // Reset the repository

	// Check the response status code
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, but got %d", http.StatusInternalServerError, w.Code)
	}
}

func TestUserAPI_GetUserByNameSuccess(t *testing.T) {
	// Create a new Gin context for testing
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	t.Setenv("DB_CONNECTIONSTRING", "something")

	// Set up the test case
	username := "whatever"
	c.AddParam(PATH_PARAM_USERNAME, username)

	// Mock the repository
	oldRepository := persistence.REPOSITORY
	persistence.REPOSITORY = persistence.GetRepository()
	persistence.REPOSITORY = MockRepository{
		Repository: oldRepository,
		err:        nil,
		user: entities.User{
			Id:        "1",
			FirstName: "Jon",
			LastName:  "Doe",
		},
	}

	// Create a new instance of UserAPI
	api := &UserAPI{}

	// Call the GetUserByName function
	api.GetUserByName(c)
	persistence.REPOSITORY = oldRepository // Reset the repository

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusInternalServerError, w.Code)
	}

	// Parse the response body
	var user User
	err := json.Unmarshal(w.Body.Bytes(), &user)
	if err != nil {
		t.Errorf("Failed to parse response body: %v", err)
	}

	// Check the user fields
	expectedUser := User{
		Id:        "1",
		FirstName: "Jon",
		LastName:  "Doe",
	}

	if !reflect.DeepEqual(user, expectedUser) {
		t.Errorf("Expected user %+v, but got %+v", expectedUser, user)
	}
}

// MockRepository is a mock implementation of the Repository interface
type MockRepository struct {
	persistence.Repository
	err  error
	user entities.User
}

func (m MockRepository) GetUserByUserName(username string) (*entities.User, error) {
	return &m.user, m.err
}
