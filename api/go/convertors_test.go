package api

import (
	"reflect"
	"testing"

	"github.com/shivinkapur/sample-go-api/persistence/entities"
)

func TestConvertNewUserToPersistenceModel(t *testing.T) {
	// Create a new user
	user := User{
		Id:         "1",
		Username:   "user1",
		FirstName:  "John",
		LastName:   "Doe",
		Email:      "john.doe@example.com",
		Password:   "password123",
		Phone:      "1234567890",
		UserStatus: 1,
	}

	// Call the convertNewUserToPersistenceModel function
	persistenceModel, err := convertNewUserToPersistenceModel(user)

	// Check for any errors
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Create the expected persistence model
	expectedPersistenceModel := &entities.User{
		Id:         "1",
		Username:   "user1",
		FirstName:  "John",
		LastName:   "Doe",
		Email:      "john.doe@example.com",
		Password:   "password123",
		Phone:      "1234567890",
		UserStatus: 1,
	}

	// Compare the persistence models
	if !reflect.DeepEqual(persistenceModel, expectedPersistenceModel) {
		t.Errorf("Expected persistence model %+v, but got %+v", expectedPersistenceModel, persistenceModel)
	}
}
