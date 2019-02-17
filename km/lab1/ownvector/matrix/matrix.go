package matrix

import (
	"math"
)

type Matrix [][]float64

func GetOwnRelativeVector(in Matrix) []float64 {
	vector := make([]float64, 0)
	for iKey, iValue := range in {
		val := 1.0
		for jKey, jValue := range iValue {
			if iKey == jKey {
				continue;
			}

			val *= jValue
		}
		vector = append(vector, math.Pow(val, 1.0/3))
	}

	return getRelativeValues(vector);
}

func getRelativeValues(vector []float64) []float64 {
	vectorSum := sum(vector)
	result := make([]float64, 0)
	for _, val := range vector {
		result = append(result, val/vectorSum)
	}

	return result;
}

func sum(vector []float64) float64 {
	sum := 0.0
	for _, val := range vector {
		sum += val
	}

	return sum;
}
