package user

import "gorm.io/gorm"

type Repository interface {
	FindUserByEmail(email string) (User, error)
	SaveUser(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepositry(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SaveUser(user User) (User, error) {
	err := r.db.Create(&user).Error

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
