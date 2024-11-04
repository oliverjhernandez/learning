package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"casita/internal/db"
	"casita/internal/validator"

	"github.com/go-chi/chi/v5"
)

type Envelope struct {
	Metadata *db.Metadata `json:"metadata,omitempty"`
	Status   string       `json:"status"`
	Message  string       `json:"message"`
	Data     interface{}  `json:"data,omitempty"`
	Error    interface{}  `json:"error,omitempty"`
}

func writeJSON(w http.ResponseWriter, data interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	return encoder.Encode(data)
}

func readJSON(r *http.Request, dst interface{}) error {
	maxBytes := 1_048_576

	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)
		case err.Error() == "http: request body too large":
			return fmt.Errorf("body must not be larger than %d bytes", maxBytes)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}

	return nil
}

func readString(r *http.Request, key string, defaultValue string) string {
	s := chi.URLParam(r, "key")

	if s == "" {
		return defaultValue
	}

	return s
}

func readInt(r *http.Request, key string, defaultValue int, v *validator.Validator) int {
	s := chi.URLParam(r, key)
	if s == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		v.AddError(key, "must be an integer value")
		return defaultValue
	}

	return i
}
