package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(string []string) ([]float64, error) {
	var floats []float64
	for _, stringVal := range string {
		floatVal, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return nil, errors.New("an error ocured while converting the strings")
		}
		floats = append(floats, floatVal)
	}

	return floats, nil
}
