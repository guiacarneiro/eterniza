package request

type BuscaFicha struct {
	TinyID string `json:"tinyId" binding:"required"`
}

type SalvaFicha struct {
	ID         uint     `json:"id"`
	TinyID     string   `json:"tinyId" binding:"required"`
	Descricao  string   `json:"descricao" binding:"required"`
	Componetes []string `json:"componetes"`
}
