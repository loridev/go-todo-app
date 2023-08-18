package utils

import (
	"errors"

	"golang.org/x/exp/slices"
)

type DBOperation string

func ValidateOperation(o DBOperation) error {
	supported_operations := []DBOperation{"show, insert, update, delete"}

	isValid := slices.Contains(supported_operations, o)

	if !isValid {
		return errors.New("invalid database operation")
	}

	return nil
}
