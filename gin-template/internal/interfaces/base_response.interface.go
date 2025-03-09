package interfaces

type BaseResponse struct {
	Code    int     `json:"code"`              // Required, non-nullable
	Message *string `json:"message,omitempty"` // Nullable string
}
