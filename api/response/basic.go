package response

type Basic struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}
