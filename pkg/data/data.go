package data

import "github.com/jinzhu/copier"

type DataService interface {
	Copy(to interface{}, from interface{}) error
	CopyWithOption(to interface{}, from interface{}, opt copier.Option) error
}

type dataService struct{}

func NewDataService() DataService {
	return &dataService{}
}

func (s *dataService) Copy(to interface{}, from interface{}) error {
	err := copier.Copy(to, from)
	return err
}

func (s *dataService) CopyWithOption(to interface{}, from interface{}, opt copier.Option) error {
	err := copier.CopyWithOption(to, from, opt)
	return err
}
