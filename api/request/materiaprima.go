package request

type CreateMateriaPrima struct {
	Label string `json:"label" binding:"required"`
	Unity int    `json:"unity" binding:"required"`
	Value string `json:"value" binding:"required"`
}
