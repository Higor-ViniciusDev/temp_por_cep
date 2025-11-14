package usecase

import (
	"github.com/Higor-ViniciusDev/temp_por_cep/internal/entity"
)

type TemperaturaPorCepUseCase struct {
	ServiceCep     entity.CepService
	ServiceWeather entity.WeatheraService
}

func NovoTemperaturaUseCase(serviceCep entity.CepService, serviceWeather entity.WeatheraService) *TemperaturaPorCepUseCase {
	return &TemperaturaPorCepUseCase{
		ServiceCep:     serviceCep,
		ServiceWeather: serviceWeather,
	}
}

type CepInputDTO struct {
	Cep string `json:"cep"`
}

type TemperaturaOutputDTO struct {
	Celsius    float64 `json:"temp_C"`
	Fahrenheit float64 `json:"temp_F"`
	Kelvin     float64 `json:"temp_K"`
}

func (u *TemperaturaPorCepUseCase) Execute(input CepInputDTO) (*TemperaturaOutputDTO, error) {
	cepEnt, err := entity.NovoCep(input.Cep)

	if err != nil {
		return nil, err
	}

	cep, err := u.ServiceCep.BuscarCepViaService(cepEnt.CEP)
	if err != nil {
		return nil, err
	}

	temp, err := u.ServiceWeather.BuscarTemperaturaPorEndereco(cep.Localidade, cep.UF)

	if err != nil {
		return nil, err
	}
	temp.ConverterCelsiusParaKelvin()

	return &TemperaturaOutputDTO{
		Fahrenheit: temp.TempFar,
		Celsius:    temp.TempCelsius,
		Kelvin:     temp.TempKelvin,
	}, nil
}
