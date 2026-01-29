package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error
	FindByID(id string) (*User, error)
	FindAll() ([]*User, error)
	Update(user *User) error
	Delete(id string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *repository) FindByID(id string) (*User, error) {
	var user User
	err := r.db.First(&user, "id = ?", id).Error
	return &user, err
}

func (r *repository) FindAll() ([]*User, error) {
	var users []*User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) Update(user *User) error {
	return r.db.Save(user).Error
}

func (r *repository) Delete(id string) error {
	return r.db.Delete(&User{}, "id = ?", id).Error
}
