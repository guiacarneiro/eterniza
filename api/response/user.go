package response

type Login struct {
	Basic
	Token string `json:"token"`
}
