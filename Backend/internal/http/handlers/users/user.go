package users

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"synergy/internal/storage"
	"synergy/internal/types"
	"synergy/internal/utils/responses"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserHandler(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateUserReq
		bool2, err := storage.CheckEmail(req.Email)

		if bool2 {
			responses.WriteJson(w, 400, responses.GeneralError(fmt.Errorf("email already exists")))
		}
		if err != nil {
			slog.Info("could not run the query")
		}
		// Decode JSON body into request struct
		err = json.NewDecoder(r.Body).Decode(&req)
		if errors.Is(err, io.EOF) {
			responses.WriteJson(w, http.StatusBadRequest, responses.GeneralError(fmt.Errorf("body is empty")))
			return
		}
		if err != nil {
			responses.WriteJson(w, http.StatusBadRequest, responses.GeneralError(err))
			return
		}

		// Validate input (only use CreateUserReq validation)
		if err := validator.New().Struct(req); err != nil {
			responses.WriteJson(w, http.StatusBadRequest, responses.ValidateError(err.(validator.ValidationErrors)))
			return
		}

		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			responses.WriteJson(w, http.StatusInternalServerError, responses.GeneralError(fmt.Errorf("failed to hash password: %w", err)))
			return
		}

		// Populate User struct
		user := &types.User{
			Name:         req.Name,
			Email:        req.Email,
			PasswordHash: string(hashedPassword),
		}

		// Create user in DB (user.ID will be populated with the SERIAL id)
		if err := storage.CreateUser(user); err != nil {
			responses.WriteJson(w, http.StatusBadRequest, responses.GeneralError(fmt.Errorf("could not insert user: %w", err)))
			return
		}

		// Respond with the generated ID (integer)
		responses.WriteJson(w, http.StatusCreated, map[string]int64{"id": user.ID})
	}
}
