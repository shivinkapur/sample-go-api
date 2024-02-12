package api

import "github.com/shivinkapur/sample-go-api/persistence/entities"

func convertNewUserToPersistenceModel(user User) (*entities.User, error) {
	persistenceModel := &entities.User{
		Id:         user.Id,
		Username:   user.Username,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		Password:   user.Password,
		Phone:      user.Phone,
		UserStatus: user.UserStatus,
	}

	return persistenceModel, nil
}

func convertPersistenceModelToUser(persistenceModel *entities.User) (*User, error) {
	user := &User{
		Id:         persistenceModel.Id,
		Username:   persistenceModel.Username,
		FirstName:  persistenceModel.FirstName,
		LastName:   persistenceModel.LastName,
		Email:      persistenceModel.Email,
		Password:   persistenceModel.Password,
		Phone:      persistenceModel.Phone,
		UserStatus: persistenceModel.UserStatus,
	}

	return user, nil
}
