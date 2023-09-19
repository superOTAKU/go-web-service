package repo

import (
	"go-web-service/pkg/model"

	"gorm.io/gorm"
)

var Users UserRepo

type UserRepo interface {
	FindAll() ([]model.User, error)
}

type userRepo struct {
	conn *gorm.DB
}

func InitUserRepo(db *gorm.DB) {
	Users = &userRepo{conn: db}
}

func (r *userRepo) FindAll() ([]model.User, error) {
	var users []model.User
	return users, r.conn.Find(&users).Error
}
