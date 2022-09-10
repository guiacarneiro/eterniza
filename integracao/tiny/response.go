package tiny

type ResponsePesquisaProdutos struct {
	Retorno struct {
		StatusProcessamento string `json:"status_processamento"`
		Status              string `json:"status"`
		Pagina              int    `json:"pagina"`
		NumeroPaginas       int    `json:"numero_paginas"`
		CodigoErro          string `json:"codigo_erro"`
		Erros               []struct {
			Erro string `json:"erro"`
		} `json:"erros"`
		Produtos []struct {
			Produto `json:"produto"`
		} `json:"produtos"`
	} `json:"retorno"`
}

type ResponseObterProduto struct {
	Retorno struct {
		StatusProcessamento string `json:"status_processamento"`
		Status              string `json:"status"`
		Pagina              int    `json:"pagina"`
		NumeroPaginas       int    `json:"numero_paginas"`
		CodigoErro          string `json:"codigo_erro"`
		Erros               []struct {
			Erro string `json:"erro"`
		} `json:"erros"`
		Produto Produto `json:"produto"`
	} `json:"retorno"`
}
