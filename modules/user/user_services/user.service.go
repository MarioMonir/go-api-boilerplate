package user_services

import "go-api-boilerplate/modules/user/user_repos"

type UserService struct {
	// logger *log.Logger
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) GetListUserService() string {
	userRepo := user_repos.NewUserRepo()

	return userRepo.GetListUserRepo()
}
