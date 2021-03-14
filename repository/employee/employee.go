package employee

import "github.com/silvergama/studying-api-with-golang/app/model"

var instance = &repositoryImpl{}

type Repository interface {
	FindById(ID string) (domain.Employee, error)
}

type repositoryImpl struct{}

func (r *repositoryImpl) FindById(ID string) (model.Employee, error) {
	result := model.Employee{}
}
