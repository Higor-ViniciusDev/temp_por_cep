package infra

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Higor-ViniciusDev/temp_por_cep/internal/entity"
	"github.com/Higor-ViniciusDev/temp_por_cep/internal/errs"
	"github.com/Higor-ViniciusDev/temp_por_cep/internal/usecase"
)

type TempHandler struct {
	TempPorCep *usecase.TemperaturaPorCepUseCase
}

func NovoTempHandler(usecase *usecase.TemperaturaPorCepUseCase) *TempHandler {
	return &TempHandler{
		TempPorCep: usecase,
	}
}

func (t *TempHandler) BuscarTemperaturaPorCep(w http.ResponseWriter, r *http.Request) {
	var input usecase.CepInputDTO

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	outPut, err := t.TempPorCep.Execute(input)

	if err != nil {
		var httpErr *errs.HttpError

		switch {
		case errors.Is(err, entity.ErrorCepInvalido):
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return

		case errors.As(err, &httpErr):
			http.Error(w, httpErr.Error(), httpErr.CodigoErro)
			return

		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(outPut)
}
