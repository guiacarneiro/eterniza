package request

type CreateUser struct {
	FisrtName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Register  bool   `json:"register"`
}
