package infra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuscarCepViaServiceValido(t *testing.T) {
	serviceCep := NovoViaCepService()
	dadosCep, err := serviceCep.BuscarCepViaService("15771034")

	assert.Nil(t, err, "Não pode haver error na busca")
	assert.NotEmpty(t, dadosCep, "Não pode ser cep vazio")
}

func TestBuscarCepViaServiceNaoEncontrado(t *testing.T) {
	serviceCep := NovoViaCepService()
	dadosCep, err := serviceCep.BuscarCepViaService("15771031")

	assert.NotNil(t, err, "Error http de cep não encontrado")
	assert.Empty(t, dadosCep, "Cep não encontrado, retornar vazio")
}
