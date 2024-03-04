package repository

import (
	"go-api/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

// IUserRepository=IUserRepositoryの全メソッドが実装されたuserRepositoryの実体
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

// structのフィールドはポインタを通してコピー先ではなく直接参照できる
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

// structのフィールドはポインタを通してコピー先ではなく直接参照できる
func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
