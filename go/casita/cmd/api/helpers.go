package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"casita/internal/models"
	"casita/internal/validator"

	"github.com/gofiber/fiber/v2"
)

type Envelope struct {
	Metadata *models.Metadata `json:"metadata,omitempty"`
	Status   string           `json:"status"`
	Message  string           `json:"message"`
	Data     interface{}      `json:"data,omitempty"`
	Error    interface{}      `json:"error,omitempty"`
}

func writeJSON(c *fiber.Ctx, status int, message string, data interface{}, metadata *models.Metadata, error string) error {
	response := Envelope{
		Metadata: metadata,
		Status:   strconv.Itoa(status),
		Message:  message,
		Data:     data,
		Error:    error,
	}

	c.Response().Header.Add("Content-Type", "application/json")

	if err := c.Status(status).JSON(response); err != nil {
		return err
	}

	return nil
}

func readJSON(c *fiber.Ctx, dst interface{}) error {
	maxBytes := 1_048_576

	err := c.BodyParser(&dst)
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

// NOTE: Not necessary when using Fiber
func readString(c *fiber.Ctx, key string, defaultValue string) string {
	s := c.Query(key, defaultValue)

	if s == "" {
		return defaultValue
	}

	return s
}

// NOTE: Not necessary when using Fiber
func readInt(c *fiber.Ctx, key string, defaultValue int, v *validator.Validator) int {
	s := c.Query(key)
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
