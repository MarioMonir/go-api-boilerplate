package user_repos

type UserRepo struct {
	// logger *log.Logger
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (u *UserRepo) GetListUserRepo() string {
	return "Hello from get the users handler from user repo !"
}
