package interfaces

type Base struct {
	Code    int     `json:"code"`              // Required, non-nullable
	Message *string `json:"message,omitempty"` // Nullable string
}
