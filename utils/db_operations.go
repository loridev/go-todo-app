package utils

import (
	"errors"

	"golang.org/x/exp/slices"
)

type DBOperation string

func ValidateOperation(o DBOperation) error {
	supportedOperations := []DBOperation{"show, insert, update, delete"}

	isValid := slices.Contains(supportedOperations, o)

	if !isValid {
		return errors.New("invalid database operation")
	}

	return nil
}
