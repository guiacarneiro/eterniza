package tiny

type RequestPesquisaProdutos struct {
	Pesquisa     string  `schema:"pesquisa,required"`
	IdTag        *int    `schema:"idTag,omitempty"`
	IdListaPreco *int    `schema:"idListaPreco,omitempty"`
	Pagina       *int    `schema:"pagina,omitempty"`
	GTIN         *string `schema:"gtin,omitempty"`
	Situacao     *string `schema:"situacao,omitempty"`
}

type RequestObterProduto struct {
	ID string `schema:"id,omitempty"`
}
