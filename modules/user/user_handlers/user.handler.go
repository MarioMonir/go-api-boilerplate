package user_handlers

import (
	"fmt"
	"go-api-boilerplate/modules/user/user_services"
	"net/http"
)

type UserHandler struct {
	// logger *log.Logger
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetListUserHandler(w http.ResponseWriter, r *http.Request) {
	userService := user_services.NewUserService()

	fmt.Fprintln(w, userService.GetListUserService())
}

func (u *UserHandler) RegisterUserHanlders(mux *http.ServeMux) {
	// Get List
	mux.HandleFunc("GET /user/", u.GetListUserHandler)
}
