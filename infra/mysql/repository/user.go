package repository

import (
	"elearning/config"
	"elearning/domain/entity"
	"elearning/domain/repository"
	"elearning/infra/mysql"
	"elearning/infra/mysql/model"
	dataPkg "elearning/pkg/data"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type userRepository struct {
	db          *mysql.Database
	dataService dataPkg.DataService
}

func NewUserRepository(db *mysql.Database, dataService dataPkg.DataService) repository.UserRepository {
	return &userRepository{
		db:          db,
		dataService: dataService,
	}
}
func (r *userRepository) Create(input *entity.User) (*entity.User, error) {
	user := &model.User{}
	err := r.dataService.Copy(user, input)
	if err != nil {
		return nil, err
	}
	result := r.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	err = r.dataService.Copy(input, user)
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (r *userRepository) Update(id int, input *entity.User) (*entity.User, error) {
	userData := &model.User{}
	if err := r.dataService.Copy(userData, input); err != nil {
		return nil, err
	}
	userData.Id = id
	query := r.db.Model(&model.User{}).
		Select("departmentname__c", "name").
		Where("id = ? ", id)
	if input.UpdatedAt.UTC().Second() > 0 {
		query.Where("recordupdatetime__c = ?", input.UpdatedAt.UTC())
	}
	query.Updates(userData)
	if err := query.Error; err != nil {
		return nil, err
	}
	if query.RowsAffected == 0 {
		return nil, gorm.ErrNotImplemented
	}
	if err := r.dataService.Copy(input, userData); err != nil {
		return nil, err
	}
	return input, nil
}

func (r *userRepository) GetByEmail(email string) (*entity.User, error) {
	user := &model.User{}
	err := r.db.Model(&model.User{}).
		Where("mail_e = ?", email).
		Where("status__c = ?", config.StatusActive).First(&user).Error
	if err != nil {
		return nil, err
	}
	result := &entity.User{}
	err = r.dataService.Copy(result, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepository) GetById(id int) (*entity.User, error) {
	user := &model.User{}
	if err := r.db.Model(&model.User{}).Where("id = ?", id).
		First(&user).Error; err != nil {
		return nil, err
	}
	result := &entity.User{}
	if err := r.dataService.Copy(result, user); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepository) GetByListId(ids []int) ([]*entity.User, error) {
	var users []*model.User
	if err := r.db.Model(&model.User{}).
		Where("id IN ?", ids).
		Where("status__c = ?", config.StatusActive).
		Find(&users).Error; err != nil {
		return nil, err
	}
	var result []*entity.User
	for _, val := range users {
		user := &entity.User{}
		if err := r.dataService.Copy(user, val); err != nil {
			return nil, err
		}
		result = append(result, user)
	}
	if len(result) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return result, nil
}

func (r *userRepository) ListByEmails(email []string) ([]*entity.User, error) {
	var users []*model.User
	if err := r.db.Model(&model.User{}).Where("mail_e IN ?", email).Find(&users).Error; err != nil {
		return nil, err
	}
	var result []*entity.User
	for _, val := range users {
		user := &entity.User{}
		if err := r.dataService.Copy(user, val); err != nil {
			return nil, err
		}
		result = append(result, user)
	}
	return result, nil
}

func (r *userRepository) UpdateLastLoginAt(id int, input *entity.User) (*entity.User, error) {
	userData := &model.User{}
	if err := r.dataService.Copy(userData, input); err != nil {
		return nil, err
	}
	userData.Id = id
	//userData.UpdatedAt = input.UpdatedAt.UTC() //TODO: replace by recordupdatetime__c
	query := r.db.Exec(fmt.Sprintf(`UPDATE "mmsp_m_herokuuser__c" SET "lastlogin_at__c"='%s' WHERE id = %v`, input.LastLoginAt.Format(time.RFC3339), input.Id))
	if err := query.Error; err != nil {
		return nil, err
	}
	if query.RowsAffected == 0 {
		return nil, gorm.ErrNotImplemented
	}
	if err := r.dataService.Copy(input, userData); err != nil {
		return nil, err
	}
	return input, nil
}

func (r *userRepository) DisableManyUser(ids []int, updateById string) error {
	if err := r.db.Model(&model.User{}).
		Where("id IN ? ", ids).
		Updates(map[string]interface{}{
			"status__c":             config.StatusInactive,
			"recordupdateuserid__c": updateById,
		}).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) ChangePassword(id int, password string) error {
	userData := &model.User{}
	if err := r.dataService.Copy(userData, entity.User{Password: password, LastChangePasswordAt: time.Now().UTC()}); err != nil {
		return err
	}
	if err := r.db.Model(&model.User{}).Where("id = ?", id).Omit("LastLoginAt").Updates(userData).Error; err != nil {
		return err
	}
	return nil
}
