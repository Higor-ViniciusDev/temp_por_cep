package main

import (
	"github.com/Higor-ViniciusDev/temp_por_cep/internal/infra"
	"github.com/Higor-ViniciusDev/temp_por_cep/internal/usecase"
)

func main() {
	useCaseTemp := usecase.NovoTemperaturaUseCase(infra.NovoViaCepService(), infra.NovoWeatherService())
	novoHandler := infra.NovoTempHandler(useCaseTemp)
	web := infra.NovoWebServer(":8000")
	web.RegistrarRota("/temperatura", novoHandler.BuscarTemperaturaPorCep, "POST")

	web.IniciarWebServer()
}
