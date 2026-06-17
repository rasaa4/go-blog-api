package helper

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetID(r *http.Request) (int, error) {

	idStr := chi.URLParam(r, "id")

	return strconv.Atoi(idStr)

}

func GetUserID(r *http.Request) (int, bool) {
	userID := r.Context().Value("user_id")

	uid, ok := userID.(float64)
	if !ok {
		return 0, false
	}
	return int(uid), true
}
