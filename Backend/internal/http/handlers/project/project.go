package project

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"synergy/internal/storage"
	"synergy/internal/types"
	"synergy/internal/utils/responses"

	"github.com/go-playground/validator/v10"
)

func CreateProjectHandler(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateProjectReq

		err := json.NewDecoder(r.Body).Decode(&req)
		if errors.Is(err, io.EOF) {
			responses.WriteJson(w, http.StatusBadRequest, responses.GeneralError(fmt.Errorf("body is empty")))
			return
		}
		if err != nil {
			responses.WriteJson(w, http.StatusBadRequest, responses.GeneralError(err))
			return
		}

		if err := validator.New().Struct(req); err != nil {
			responses.WriteJson(w, http.StatusBadRequest, responses.ValidateError(err.(validator.ValidationErrors)))
			return
		}

		project := &types.Project{
			Name:    req.Name,
			OwnerID: req.CreatedBy,
		}

		if err := storage.CreateProject(project); err != nil {
			responses.WriteJson(w, http.StatusBadRequest, responses.GeneralError(fmt.Errorf("could not insert project: %w", err)))
			return
		}

		responses.WriteJson(w, http.StatusCreated, map[string]interface{}{"id": project.ID})
	}
}

func GetProjectByIDHandler(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		if idStr == "" {
			responses.WriteJson(w, http.StatusBadRequest,
				responses.GeneralError(fmt.Errorf("missing project id")))
			return
		}

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			responses.WriteJson(w, http.StatusBadRequest,
				responses.GeneralError(fmt.Errorf("invalid project id: %w", err)))
			return
		}

		proj, err := storage.GetProjectByID(id)
		if err != nil {
			responses.WriteJson(w, http.StatusNotFound,
				responses.GeneralError(err))
			return
		}

		responses.WriteJson(w, http.StatusOK, proj)
	}
}
