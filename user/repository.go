package user

import "gorm.io/gorm"

type Repository interface {
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
