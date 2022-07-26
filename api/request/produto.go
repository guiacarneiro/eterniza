package request

type BuscaFicha struct {
	TinyID string `json:"tinyId" binding:"required"`
}
