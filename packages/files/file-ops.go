package files

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to find file.")
	}

	fileString := string(data)
	value, err := strconv.ParseFloat(fileString, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored value.")
	}

	return value, nil
}

func SetFloatInFile(fileName string, number float64) {
	numberText := fmt.Sprint(number)
	os.WriteFile(fileName, []byte(numberText), 0644)
}
