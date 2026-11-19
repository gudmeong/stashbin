package response

type ErrorResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func NewError(msg string) *ErrorResponse {
	return &ErrorResponse{Ok: false, Message: msg}
}

var (
	ErrInternalServerError = NewError("internal server error")
	ErrDocumentNotFound    = NewError("document not found")
	ErrInvalidBody         = NewError("invalid body")
	ErrContentEmpty        = NewError("content is empty")
	ErrNoKeyQuery          = NewError("no document key provided")
)
