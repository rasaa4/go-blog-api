package handler

import (
	"blog-api/internal/domain"
	"blog-api/internal/dto"
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

	var req dto.RegisterUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body")
		return
	}

	if err := req.Validate(); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	user := domain.User{
		Username: req.Username,
		Password: req.Password,
	}
	createdUser, err := h.usecase.Register(user)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(w, createdUser)

}
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	var req dto.LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body")
		return
	}
	if err := req.Validate(); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return

	}
	token, err := h.usecase.Login(
		req.Username, req.Password)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, "invalid credentials")
		return
	}

	response.Success(w, map[string]string{

		"token": token,
	})
}
