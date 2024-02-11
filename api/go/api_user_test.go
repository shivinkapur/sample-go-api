package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUserAPI_GetUserByName(t *testing.T) {
	// Create a new Gin context for testing
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Create a new instance of UserAPI
	api := &UserAPI{}

	// Call the GetUserByName function
	api.GetUserByName(c)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	// Parse the response body
	var user User
	err := json.Unmarshal(w.Body.Bytes(), &user)
	if err != nil {
		t.Errorf("Failed to parse response body: %v", err)
	}

	// Check the user fields
	expectedUser := User{
		Id:         1,
		Username:   "user1",
		FirstName:  "John",
		LastName:   "Doe",
		Email:      "johndoe1@ok.com",
		Password:   "123456",
		Phone:      "1234567890",
		UserStatus: 1,
	}

	if !reflect.DeepEqual(user, expectedUser) {
		t.Errorf("Expected user %+v, but got %+v", expectedUser, user)
	}
}
