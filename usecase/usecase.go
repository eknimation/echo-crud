package usecase

import (
	"echo-crud/repositories"
)

type Usecase struct {
	Repositories *repositories.Repositiries
}

func New(r *repositories.Repositiries) *Usecase {
	return &Usecase{
		Repositories: r,
	}
}

func (u *Usecase) U1() string {
	result, err := u.Repositories.GetName(1)
	if err != nil {
		return err.Error()
	}
	return result
}
