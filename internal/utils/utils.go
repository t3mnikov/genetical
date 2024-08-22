package utils

// Функция для ограничения значений внутри диапазона
func FixVal(val, min, max int) int {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}
