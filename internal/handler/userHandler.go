package handler

import (
	"blog-api/internal/domain"
	"blog-api/internal/dto"
	"blog-api/internal/response"
	"blog-api/internal/richerror"
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
		rErr := richerror.New("User.Register").
			WithKind(richerror.KindInvalid).
			WithMessage("invalid request body").
			WithError(err)

		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return
	}

	if err := req.Validate(); err != nil {
		rErr := richerror.New("User.Register").
			WithKind(richerror.KindInvalid).
			WithMessage(err.Error()).
			WithError(err)

		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return
	}

	user := domain.User{
		Username: req.Username,
		Password: req.Password,
	}

	createdUser, err := h.usecase.Register(user)
	if err != nil {
		rErr := richerror.New("User.Register").
			WithKind(richerror.KindUnexpected).
			WithMessage("failed to register user").
			WithError(err)

		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return
	}

	response.Success(w, createdUser)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	var req dto.LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		rErr := richerror.New("User.Login").
			WithKind(richerror.KindInvalid).
			WithMessage("invalid request body").
			WithError(err)
		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return
	}
	if err := req.Validate(); err != nil {
		rErr := richerror.New("User.Login").
			WithKind(richerror.KindInvalid).
			WithMessage(err.Error()).
			WithError(err)
		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return

	}
	token, err := h.usecase.Login(
		req.Username, req.Password)
	if err != nil {
		rErr := richerror.New("User.Login").
			WithKind(richerror.KindForbidden).
			WithMessage("invalid credentials").
			WithError(err)
		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return
	}

	response.Success(w, map[string]string{

		"token": token,
	})
}
