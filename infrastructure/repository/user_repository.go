package repository

import (
	"clean-architecture/domain/model"
	"clean-architecture/infrastructure/db"
)

type UserRepository interface {
	Create(username string, hashedPassword string) (*model.User, error)
}

type repository struct {
	*db.Database
}

func NewUserRepository(db *db.Database) UserRepository {
	var userRepository UserRepository
	userRepository = &repository{db}
	return userRepository
}

func (r *repository) Create(username string, hashedPassword string) (*model.User, error) {
	user := &model.User{
		UserName: username,
		Password: hashedPassword,
	}
	// fmt.Printf("%+v\n", r.DB)
	result := r.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
