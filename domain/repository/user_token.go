package repository

import "elearning/domain/entity"

type UserTokenRepository interface {
	GetByToken(token string) (*entity.UserToken, error)
	GetByTokenForEmail(token string) (*entity.UserToken, error)
	GetByUserId(id int) (*entity.UserToken, error)
	Create(input *entity.UserToken) (*entity.UserToken, error)
	CreateMany(input []*entity.UserToken) ([]*entity.UserToken, error)
	Delete(id int) error
}
