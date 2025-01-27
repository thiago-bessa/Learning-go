package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(filename string) (float64, error) {
	var data, readErr = os.ReadFile(filename)

	if readErr != nil {
		return 0, errors.New("Failed to read file.")
	}

	var floatValueText = string(data)
	var floatValue, parseErr = strconv.ParseFloat(floatValueText, 64)

	if parseErr != nil {
		return 0, errors.New("Failed to parse stored value.")
	}

	return floatValue, nil
}

func WriteFloatToFile(floatValue float64, filename string) {
	var floatValueText = fmt.Sprint(floatValue)
	os.WriteFile(filename, []byte(floatValueText), 0644)
}
