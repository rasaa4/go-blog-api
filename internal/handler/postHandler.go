package handler

import (
	"blog-api/internal/domain"
	"blog-api/internal/dto"
	"blog-api/internal/response"
	"blog-api/internal/richerror"
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

	var req dto.CreatePostRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		rErr := richerror.New("Post.Create").
			WithKind(richerror.KindInvalid).
			WithMessage("invalid request body").
			WithError(err)
		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return
	}
	if err := req.Validate(); err != nil {
		rErr := richerror.New("Post.Create").
			WithKind(richerror.KindInvalid).
			WithMessage(err.Error()).
			WithError(err)
		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return
	}
	userID := r.Context().Value("user_id")
	uid, ok := userID.(float64)
	if !ok {
		rErr := richerror.New("Post.Create").
			WithKind(richerror.KindUnauthorized).
			WithMessage("invalid user").
			WithError(nil)
		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return
	}
	post := domain.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  int(uid),
	}

	posted, err := h.usecase.Create(post)
	if err != nil {
		rErr := richerror.New("Post.Create").
			WithKind(richerror.KindUnexpected).
			WithMessage("failed to create post").
			WithError(err)
		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return
	}
	response.Success(w, posted)
}
func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {

	posts, err := h.usecase.GetPosts()
	if err != nil {
		rErr := richerror.New("Post.GetPosts").
			WithKind(richerror.KindUnexpected).
			WithMessage("failed to fetch posts").
			WithError(err)
		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return
	}
	response.Success(w, posts)

}

func (h *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		rErr := richerror.New("Post.GetByID").
			WithKind(richerror.KindInvalid).
			WithMessage("invalid post id").
			WithError(err)
		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return
	}
	post, err := h.usecase.GetByID(id)
	if err != nil {
		rErr := richerror.New("Post.GetByID").
			WithKind(richerror.KindNotFound).
			WithMessage("post not found").
			WithError(err)
		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return
	}

	response.Success(w, post)
}
func (h *PostHandler) Delete(w http.ResponseWriter, r *http.Request) {

	//id from path
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		rErr := richerror.New("Post.Delete").
			WithKind(richerror.KindInvalid).
			WithMessage("invalid id").
			WithError(err)
		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return
	}
	//id from jwt
	userID := r.Context().Value("user_id")
	uid, ok := userID.(float64)
	if !ok {
		rErr := richerror.New("Post.Delete").
			WithKind(richerror.KindUnauthorized).
			WithMessage("invalid user").
			WithError(nil)

		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return
	}
	post, err := h.usecase.GetByID(id)
	if err != nil {
		rErr := richerror.New("Post.Delete").
			WithKind(richerror.KindNotFound).
			WithMessage("post not found").
			WithError(err)
		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return
	}
	if post.UserID != int(uid) {
		rErr := richerror.New("Post.Delete").
			WithKind(richerror.KindForbidden).
			WithMessage("you are not allowed to delete this post").
			WithError(nil)
		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return
	}
	err = h.usecase.Delete(id)
	if err != nil {
		rErr := richerror.New("Post.Delete").
			WithKind(richerror.KindUnexpected).
			WithMessage("failed to delete post").
			WithError(err)
		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
func (h *PostHandler) Update(w http.ResponseWriter, r *http.Request) {

	//گرفتن ایدی از url

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		rErr := richerror.New("Post.Update").
			WithKind(richerror.KindInvalid).
			WithMessage("invalid id").
			WithError(err)
		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return
	}
	//گرفتن یوزر ایدی از jwt
	userID := r.Context().Value("user_id")
	uid, ok := userID.(float64)
	if !ok {
		rErr := richerror.New("Post.Update").
			WithKind(richerror.KindUnauthorized).
			WithMessage("invalid user").
			WithError(nil)
		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return
	}
	existingPost, err := h.usecase.GetByID(id)
	if err != nil {
		rErr := richerror.New("Post.Update").
			WithKind(richerror.KindNotFound).
			WithMessage("post not found").
			WithError(err)
		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return
	}
	if existingPost.UserID != int(uid) {
		rErr := richerror.New("Post.Update").
			WithKind(richerror.KindForbidden).
			WithMessage("you are not allowed to edit this post").
			WithError(nil)
		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return
	}

	//decoding body
	var req dto.UpdatePostRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		rErr := richerror.New("Post.Update").
			WithKind(richerror.KindInvalid).
			WithMessage("invalid request body").
			WithError(err)
		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return
	}
	if err := req.Validate(); err != nil {
		rErr := richerror.New("Post.Update").
			WithKind(richerror.KindInvalid).
			WithMessage(err.Error()).
			WithError(err)
		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return
	}
	post := domain.Post{
		ID:      id,
		Title:   req.Title,
		Content: req.Content,
		UserID:  existingPost.UserID,
	}
	err = h.usecase.Update(post)
	if err != nil {
		rErr := richerror.New("Post.Update").
			WithKind(richerror.KindUnexpected).
			WithMessage("failed to update post").
			WithError(err)
		response.Error(w, rErr.Kind().HTTPStatus(), rErr.Error())
		return

	}
	response.Success(w, post)

}
