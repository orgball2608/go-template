package repository

import (
	"elearning/domain/entity"
	"elearning/domain/repository"
	"elearning/infra/mysql"
	"elearning/infra/mysql/model"
	dataPkg "elearning/pkg/data"
)

type userTokenRepository struct {
	db          *mysql.Database
	dataService dataPkg.DataService
}

func NewUserTokenRepository(db *mysql.Database, dataService dataPkg.DataService) repository.UserTokenRepository {
	return &userTokenRepository{
		db:          db,
		dataService: dataService,
	}
}

func (r *userTokenRepository) GetByTokenForEmail(token string) (*entity.UserToken, error) {
	userToken := &model.UserToken{}
	if err := r.db.Model(&model.UserToken{}).
		Where("token = ?", token).
		Where("expire_at IS NULL").
		First(&userToken).Error; err != nil {
		return nil, err
	}
	result := &entity.UserToken{}
	if err := r.dataService.Copy(result, userToken); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userTokenRepository) GetByToken(token string) (*entity.UserToken, error) {
	userToken := &model.UserToken{}
	if err := r.db.Model(&model.UserToken{}).
		Where("token = ?", token).
		First(&userToken).Error; err != nil {
		return nil, err
	}
	result := &entity.UserToken{}
	if err := r.dataService.Copy(result, userToken); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userTokenRepository) Create(input *entity.UserToken) (*entity.UserToken, error) {
	userToken := &model.UserToken{}
	if err := r.dataService.Copy(userToken, input); err != nil {
		return nil, err
	}
	result := r.db.Create(userToken)
	if result.Error != nil {
		return nil, result.Error
	}
	if err := r.dataService.Copy(input, result); err != nil {
		return nil, err
	}
	return input, nil
}

func (r *userTokenRepository) CreateMany(input []*entity.UserToken) ([]*entity.UserToken, error) {
	userToken := []*model.UserToken{}
	for _, val := range input {
		tmpUserToken := &model.UserToken{}
		if err := r.dataService.Copy(tmpUserToken, val); err != nil {
			return nil, err
		}
		userToken = append(userToken, tmpUserToken)
	}
	result := r.db.Create(userToken)
	if result.Error != nil {
		return nil, result.Error
	}
	output := []*entity.UserToken{}
	for _, val := range userToken {
		tmpUserToken := &entity.UserToken{}
		if err := r.dataService.Copy(tmpUserToken, val); err != nil {
			return nil, err
		}
		output = append(output, tmpUserToken)
	}
	return output, nil
}

func (r *userTokenRepository) GetByUserId(id int) (*entity.UserToken, error) {
	userToken := &model.UserToken{}
	if err := r.db.Model(&model.UserToken{}).
		Where("user_id = ?", id).
		First(&userToken).Error; err != nil {
		return nil, err
	}
	result := &entity.UserToken{}
	if err := r.dataService.Copy(result, userToken); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userTokenRepository) Delete(id int) (err error) {
	result := r.db.Delete(&model.UserToken{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
