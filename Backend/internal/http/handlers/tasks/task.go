package tasks

import (
	"encoding/json"
	"net/http"
	"synergy/internal/storage"
	"synergy/internal/types"
	"synergy/internal/utils/responses"

	"github.com/go-playground/validator/v10"
)

func CreateTaskHandler(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task types.Task

		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			responses.WriteJson(w, http.StatusBadRequest, responses.GeneralError(err))
			return
		}

		if err := validator.New().Struct(task); err != nil {
			responses.WriteJson(w, http.StatusBadRequest, responses.ValidateError(err.(validator.ValidationErrors)))
			return
		}

		if err := storage.CreateTask(&task); err != nil {
			responses.WriteJson(w, http.StatusBadRequest, responses.GeneralError(err))
			return
		}

		responses.WriteJson(w, http.StatusCreated, task)
	}
}
