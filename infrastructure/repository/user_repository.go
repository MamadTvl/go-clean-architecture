package repository

import (
	"clean-architecture/domain/model"
	"clean-architecture/infrastructure/db"
)

type UserRepository struct {
	*db.Database
}

func NewUserRepository(db *db.Database) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(username string, hashedPassword string) (*model.User, error) {
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
