package infra

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/Higor-ViniciusDev/temp_por_cep/internal/entity"
	"github.com/valyala/fastjson"
)

type WeatherService struct{}

func NovoWeatherService() *WeatherService {
	return &WeatherService{}
}

func (s *WeatherService) BuscarTemperaturaPorEndereco(cidade string, estadoSigla string) (*entity.Temperatura, error) {
	cidadeFormatada := url.QueryEscape(cidade)

	resp, err := http.Get(fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=cd5f13c67b234405ab1151712251311&q=%v,%v,brazil&aqi=no", cidadeFormatada, estadoSigla))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var p fastjson.Parser
	data, err := p.ParseBytes(body)

	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	return &entity.Temperatura{
		TempCelsius: data.Get("current").GetFloat64("temp_c"),
		TempFar:     data.Get("current").GetFloat64("temp_f"),
	}, nil
}
