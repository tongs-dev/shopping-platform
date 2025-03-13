package common

import (
	"encoding/json"
	"errors"
	"reflect"
)

// SwapTo assigns values from `request` struct to `target` struct using JSON tags.
func SwapTo(request, target interface{}) error {
	// Validate input parameters
	if request == nil || target == nil {
		return errors.New("request or target cannot be nil")
	}

	// Ensure target is a pointer (json.Unmarshal requires a pointer)
	if reflect.TypeOf(target).Kind() != reflect.Ptr {
		return errors.New("target must be a pointer")
	}

	// Convert request struct to JSON bytes
	dataByte, err := json.Marshal(request)
	if err != nil {
		return err
	}

	// Convert JSON bytes to target struct
	err = json.Unmarshal(dataByte, target)
	if err != nil {
		return err
	}

	return nil
}
