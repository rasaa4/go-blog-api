package handler

import (
	"blog-api/internal/domain"
	"blog-api/internal/response"
	"blog-api/internal/usecase"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PostHandler struct {
	usecase *usecase.PostUsecase
}

func NewPostHandler(usecase *usecase.PostUsecase) *PostHandler {
	return &PostHandler{usecase: usecase}
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {

	var post domain.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body")

		return
	}
	userID := r.Context().Value("user_id")
	uid, ok := userID.(float64)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "invalid user")
		return
	}
	post.UserID = int(uid)

	posted, err := h.usecase.Create(post)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(w, posted)
}
func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {

	posts, err := h.usecase.GetPosts()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}
	response.Success(w, posts)

}

func (h *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "id is not valid")
		return
	}
	post, err := h.usecase.GetByID(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, "post not found")

		return
	}

	response.Success(w, post)
}
func (h *PostHandler) Delete(w http.ResponseWriter, r *http.Request) {

	//id from path
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid id")
		return
	}
	//id from jwt
	userID := r.Context().Value("user_id")
	uid, ok := userID.(float64)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "invalid user")
		return
	}
	post, err := h.usecase.GetByID(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, "post not found")
		return
	}
	if post.UserID != int(uid) {
		response.Error(w, http.StatusForbidden, "you are not allowed to delete this post")
		return
	}
	err = h.usecase.Delete(id)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
func (h *PostHandler) Update(w http.ResponseWriter, r *http.Request) {

	//گرفتن ایدی از url

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid id")
		return
	}
	//گرفتن یوزر ایدی از jwt
	userID := r.Context().Value("user_id")
	uid, ok := userID.(float64)
	if !ok {
		response.Error(w, http.StatusUnauthorized, "invalid user")
		return
	}
	existingPost, err := h.usecase.GetByID(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, "post not found")
		return
	}
	if existingPost.UserID != int(uid) {
		response.Error(w, http.StatusForbidden, "you are not allowed to edit this post")

		return
	}

	//decoding body
	var post domain.Post
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid body")

		return
	}
	post.ID = id
	err = h.usecase.Update(post)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}
	response.Success(w, post)

}
