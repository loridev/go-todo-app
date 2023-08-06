package utils

import "fmt"

func DisplayError(message string, err error) string {
	return fmt.Sprintf("%s: %s", message, err)
}
