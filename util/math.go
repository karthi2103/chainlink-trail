package util

import (
	"fmt"
	"github.com/shopspring/decimal"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
	"math"
	"math/big"
)

type Number interface {
	constraints.Float | constraints.Integer
}

// ToNormalizedFloat : convert smart contract answer to decimal
func ToNormalizedFloat(val interface{}, normalizationFactor int) (float64, bool) {
	value := new(big.Int)
	switch v := val.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(normalizationFactor)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result.Float64()
}

// Median : calculate median for given slice of numbers
func Median[T Number](data []T) float64 {
	dataCopy := make([]T, len(data))
	copy(dataCopy, data)

	slices.Sort(dataCopy)

	var median float64
	l := len(dataCopy)
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		median = float64((dataCopy[l/2-1] + dataCopy[l/2]) / 2.0)
	} else {
		median = float64(dataCopy[l/2])
	}

	return median
}

// MeasureDeviation : verifies if the individual answer from oracles are within the
// specified boundary defined by deviation percentage param
func MeasureDeviation(feed []float64, median float64, deviation float64) bool {
	deviationBoundary := (median * deviation) / 100
	for i, f := range feed {
		if math.Abs(f-median) > deviationBoundary {
			fmt.Println("failed at index: ", i)
			return false
		}
	}
	return true
}
