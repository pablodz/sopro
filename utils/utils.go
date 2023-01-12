package utils

func LimitsSliceFloat64(items []float64) (float64, float64) {
	min := items[0]
	max := items[0]
	for _, val := range items {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return min, max
}
