package numericUtils

func ChangeFloatValueInPercentage(oldValue, newValue float64) float64 {
	if oldValue == 0.0 {
		return oldValue
	}

	return (newValue - oldValue) / oldValue
}
