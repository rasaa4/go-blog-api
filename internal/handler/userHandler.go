package handler

import (
	"blog-api/internal/domain"
	"blog-api/internal/response"
	"blog-api/internal/usecase"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	usecase *usecase.UserUsecase
}

func NewUserHandler(usecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {

	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body")
		return
	}
	createdUser, err := h.usecase.Register(user)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(w, createdUser)

}
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body")
		return
	}
	token, err := h.usecase.Login(user.Username, user.Password)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	response.Success(w, map[string]string{

		"token": token,
	})
}
