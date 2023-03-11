package repository

import (
	"database/pkg/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByEmail(email string) (*models.User, error)
}

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (userStore *UserStore) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := userStore.db.Where("email=?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
