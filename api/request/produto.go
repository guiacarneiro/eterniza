package request

type BuscaFicha struct {
	TinyID string `json:"tinyId" binding:"required"`
}

type SalvaFicha struct {
	ID          uint     `json:"id"`
	TinyID      string   `json:"tinyId" binding:"required"`
	Variacao    string   `json:"variacao" binding:"required"`
	Componentes []string `json:"componentes" binding:"required"`
}
