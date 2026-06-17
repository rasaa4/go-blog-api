package httpmsg

import (
	"blog-api/internal/richerror"
	"net/http"
)

func Error(err error) (string, int) {
	switch e := err.(type) {
	case richerror.RichError:
		return e.Error(), mapKindToStatus(e.Kind())
	default:
		return "something went wrong", http.StatusInternalServerError
	}

}
func mapKindToStatus(kind richerror.Kind) int {
	switch kind {
	case richerror.KindInvalid:
		return http.StatusBadRequest
	case richerror.KindForbidden:
		return http.StatusForbidden
	case richerror.KindNotFound:
		return http.StatusNotFound
	case richerror.KindUnexpected:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
