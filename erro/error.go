package erro

import (
	"errors"
	"github.com/guiacarneiro/eterniza/logger"
)

func ERROR(e error) error {
	if !errors.Is(e, SenhaNaoEncontrada) &&
		!errors.Is(e, SuspensaoSolicitada) &&
		!errors.Is(e, LogoffSolicitado) &&
		!errors.Is(e, UsuarioNaoLogadoEmGuiche) &&
		!errors.Is(e, UsuarioNaoEncontrado) &&
		!errors.Is(e, NaoAutorizado) &&
		!errors.Is(e, SenhaObrigatoria) &&
		!errors.Is(e, LoginObrigatorio) &&
		!errors.Is(e, MetodoDesconhecido) &&
		!errors.Is(e, ServicoObrigatorio) &&
		!errors.Is(e, SemPermissao) &&
		!errors.Is(e, SenhaNaoInformada) &&
		!errors.Is(e, GuicheOcupado) &&
		!errors.Is(e, ComentarioNaoInformado) &&
		!errors.Is(e, DadosIncorretos) &&
		!errors.Is(e, ModoOffilineDesabilitado) &&
		!errors.Is(e, CadastroNaoEncontrado) &&
		!errors.Is(e, IntervaloMinimo) &&
		!errors.Is(e, MaximoRechamadas) &&
		!errors.Is(e, UsuarioNaoIdentificado) &&
		!errors.Is(e, StatusIncompativel) &&
		!errors.Is(e, ParametrosIncorretos) {
		logger.LogInfo.Println(e)
		return Desconhecido
	}
	return e
}

var (
	Desconhecido             = errors.New("Erro não identificado")
	SenhaNaoEncontrada       = errors.New("Senha não encontrada")
	SuspensaoSolicitada      = errors.New("Suspensão Solicitada")
	LogoffSolicitado         = errors.New("Logoff Solicitado")
	UsuarioNaoLogadoEmGuiche = errors.New("Atendente não logado em guichê")
	UsuarioNaoEncontrado     = errors.New("Usuário não encontrado")
	NaoAutorizado            = errors.New("Não autorizado")
	SenhaObrigatoria         = errors.New("Senha obrigatória")
	LoginObrigatorio         = errors.New("Login obrigatório")
	MetodoDesconhecido       = errors.New("Método desconhecido")
	ServicoObrigatorio       = errors.New("Pelo menos um serviço obrigatório")
	SemPermissao             = errors.New("Sem permissão")
	SenhaNaoInformada        = errors.New("Senha não informada")
	GuicheOcupado            = errors.New("Guichê ocupado")
	ComentarioNaoInformado   = errors.New("Comentário não informado")
	DadosIncorretos          = errors.New("Dados incorretos")
	ModoOffilineDesabilitado = errors.New("Modo offiline desabilitado")
	CadastroNaoEncontrado    = errors.New("Cadastro não encontrado")
	IntervaloMinimo          = errors.New("Intervalo mínimo entre chamadas não atingido")
	MaximoRechamadas         = errors.New("Máximo de rechamadas atingido")
	UsuarioNaoIdentificado   = errors.New("Usuário não identificado, realize o login")
	StatusIncompativel       = errors.New("Status incompatível")
	ParametrosIncorretos     = errors.New("Parâmetros incorretos")
)
