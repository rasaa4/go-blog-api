package richerror

type Kind int

const (
	KindInvalid Kind = iota + 1
	KindForbidden
	KindNotFound
	KindUnexpected
	KindUnauthorized
)

func (k Kind) HTTPStatus() int {
	switch k {
	case KindInvalid:
		return 400
	case KindUnauthorized:
		return 401
	case KindForbidden:
		return 403
	case KindNotFound:
		return 404
	case KindUnexpected:
		return 500
	default:
		return 500

	}
}

type Op string
type RichError struct {
	operation    Op
	wrappedError error
	message      string
	kind         Kind
	meta         map[string]interface{}
}

//constructor

func New(op Op) RichError {
	return RichError{
		operation: op,
	}
}
func (r RichError) Kind() Kind {
	return r.kind
}

func (r RichError) Err() error {
	return r.wrappedError
}

func (r RichError) Unwrap() error {
	return r.wrappedError
}
func (r RichError) Error() string {
	return r.message
}
func (r RichError) WithMessage(message string) RichError {
	r.message = message
	return r
}
func (r RichError) WithKind(kind Kind) RichError {
	r.kind = kind
	return r
}
func (r RichError) WithError(err error) RichError {
	r.wrappedError = err
	return r
}
func (r RichError) WithMeta(meta map[string]interface{}) RichError {
	r.meta = meta
	return r
}
