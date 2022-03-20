package generic

type Number interface {
	int64 | ~float64 | ~int
}

type IntType int

// SumIntOrFloat sum of map
func SumIntOrFloat[K comparable, V int | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumNumber sum of map
func SumNumber[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
