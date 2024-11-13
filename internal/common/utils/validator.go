package utils

import (
	"blockhouse_streaming_api/pkg/validator"
	"sync"
)

var lock = &sync.Mutex{}

var validatorInstance validator.Validator

func GetValidator() validator.Validator {
	if validatorInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if validatorInstance == nil {
			validatorInstance = validator.NewValidator()
		}
	}

	return validatorInstance
}
