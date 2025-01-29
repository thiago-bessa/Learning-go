package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {

	result := make([]float64, 0)

	for _, stringValue := range strings {
		floatPrice, err := strconv.ParseFloat(stringValue, 64)

		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}
		result = append(result, floatPrice)
	}

	return result, nil
}
