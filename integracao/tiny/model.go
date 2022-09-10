package tiny

type Produto struct {
	ID               string  `json:"id"`
	Codigo           string  `json:"codigo"`
	Nome             string  `json:"nome"`
	Preco            float64 `json:"preco"`
	PrecoPromocional float64 `json:"preco_promocional"`
	PrecoCusto       float64 `json:"preco_custo"`
	PrecoCustoMedio  float64 `json:"preco_custo_medio"`
	Unidade          string  `json:"unidade"`
	GTIN             string  `json:"gtin"`
	TipoVariacao     string  `json:"tipoVariacao"`
	Localizacao      string  `json:"localizacao"`
	Situacao         string  `json:"situacao"`
}
