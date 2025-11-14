package infra

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Higor-ViniciusDev/temp_por_cep/internal/entity"
	"github.com/Higor-ViniciusDev/temp_por_cep/internal/util"
)

type ViaCepService struct{}

func NovoViaCepService() *ViaCepService {
	return &ViaCepService{}
}

func (s *ViaCepService) BuscarCepViaService(cep string) (*entity.Cep, error) {
	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Cep        string `json:"cep"`
		Logradouro string `json:"logradouro"`
		Bairro     string `json:"bairro"`
		Localidade string `json:"localidade"`
		Uf         string `json:"uf"`
		Estado     string `json:"estado"`
		Erro       any    `json:"erro"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if data.Erro != nil {
		return nil, util.CepNaoEncontrado(nil)
	}

	return &entity.Cep{
		CEP:        data.Cep,
		Logradouro: data.Logradouro,
		Bairro:     data.Bairro,
		Localidade: data.Localidade,
		UF:         data.Uf,
		Estado:     data.Estado,
	}, nil
}
