package service

import (
	"rocs/model"
	"rocs/repositories"

	"github.com/mlogclub/simple"
)

type UserService struct {
	UserRepository  *repositories.UserRepository
	GithubRepostory *repositories.GithubUserRepository
}

func NewUserService() *UserService {
	return &UserService{
		UserRepository:  repositories.NewUserRepository(),
		GithubRepostory: repositories.NewGithubUserRepository(),
	}
}

func (r *UserService) Get(id int64) *model.User {
	return r.UserRepository.Get(simple.GetDB(), id)
}
