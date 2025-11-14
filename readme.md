Este projeto expõe uma API em Go que recebe um **CEP** e retorna a **temperatura atual** da cidade correspondente, realizando consultas externas ao ViaCEP e ao WeatherAPI.

A aplicação pode ser executada localmente ou via **Docker**.

---

## Executando Localmente com Docker

### 1. Build da imagem

```bash
docker build -t temperatura-api .

### 2. Executar o container

```bash
docker run -p 8000:8000 temperatura-api

## Endpoint Principal

### `POST /temperatura`

Recebe um JSON contendo o CEP:

#### Request body
```json
{
  "cep": "15771034"
}
```

### Dentro da Pasta API contém teste http para api local e web