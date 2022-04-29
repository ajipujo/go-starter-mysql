package user

import (
	"gostart/config"

	"gorm.io/gorm"
)

type Repository interface {
	FindUserByID(userID int) (User, error)
	FindUserByEmail(email string) (User, error)
	Save(user User) (User, error)
	Update(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository() *repository {
	db := config.DB()
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindUserByEmail(email string) (User, error) {
	user := User{}

	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, err
}

func (r *repository) FindUserByID(userID int) (User, error) {
	user := User{}
	err := r.db.Where("id = ?", userID).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
