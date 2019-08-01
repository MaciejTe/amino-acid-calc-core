package validators

import (
	"errors"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func ValidatePort(input string) error {
	_, err := strconv.Atoi(input)
	if err != nil {
		log.Error("REST API port validation failed")
		return errors.New("invalid port number")
	}
	return nil
}
