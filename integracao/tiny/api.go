package tiny

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/guiacarneiro/eterniza/config"
	"github.com/guiacarneiro/eterniza/util/schema"
	"io/ioutil"
	"net/http"
	"net/url"
)

var TokenTiny string
var encoder = schema.NewEncoder()

const PesquisaProdutos = "produtos.pesquisa"
const ObterProduto = "produto.obter"

func init() {
	TokenTiny = config.GetPropriedade("TOKEN_TINY")
}

func ChamaMetodoTiny(param, resposta interface{}, metodo string) error {

	urlTiny := config.GetPropriedade("TINY_URL") + metodo + ".php"

	parametro := url.Values{}
	var err error
	if param != nil {
		err = encoder.Encode(param, parametro)
		if err != nil {
			return err
		}
	}

	parametro.Add("token", TokenTiny)
	parametro.Add("formato", "json")

	client := new(http.Client)
	resp, err := client.PostForm(urlTiny, parametro)

	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("Erro resposta: %d, %s", resp.StatusCode, urlTiny))
	}
	if resposta != nil {
		b, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		err = json.Unmarshal(b, resposta)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}
