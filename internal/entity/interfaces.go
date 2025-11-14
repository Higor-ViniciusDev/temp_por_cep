package entity

type CepService interface {
	BuscarCepViaService(cep string) (*Cep, error)
}

type WeatheraService interface {
	BuscarTemperaturaPorEndereco(cidade string, estadoSigla string) (*Temperatura, error)
}
