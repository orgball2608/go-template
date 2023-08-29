package repository

import "elearning/domain/entity"

type UserRepository interface {
	GetByEmail(email string) (*entity.User, error)
	GetById(id int) (*entity.User, error)
	GetByListId(ids []int) ([]*entity.User, error)
	ListByEmails(email []string) ([]*entity.User, error)
	Update(id int, user *entity.User) (*entity.User, error)
	UpdateLastLoginAt(id int, input *entity.User) (*entity.User, error)
	Create(user *entity.User) (*entity.User, error)
	ChangePassword(id int, password string) error
	DisableManyUser(ids []int, updateById string) error
}
