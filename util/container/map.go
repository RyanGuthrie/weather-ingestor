package container

// Values takes a map and returns a slice of its values in the same order as they appear in the map.
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))

	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}

	return values
}
