package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

type CreatedPostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
type PostResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
}

func (r CreatedPostRequest) Validate() error {

	return validation.ValidateStruct(&r, validation.Field(&r.Title, validation.Required, validation.Length(3, 100)),
		validation.Field(&r.Content, validation.Required, validation.Length(10, 5000)))
}
func (r UpdatePostRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Title,
			validation.Required,
			validation.Length(3, 100),
		),
		validation.Field(&r.Content,
			validation.Required,
			validation.Length(10, 5000),
		),
	)

}
