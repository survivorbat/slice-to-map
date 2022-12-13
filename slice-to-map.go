package slicetomap

// ToMap converts a slice to a map with the keyFunc determining what the key of a value should be.
// Will override any double values.
func ToMap[T any, K comparable](slice []T, keyFunc func(T) K) map[K]T {
	m := make(map[K]T, len(slice))
	for _, v := range slice {
		m[keyFunc(v)] = v
	}
	return m
}

// ToSliceMap converts a slice to a map with the keyFunc determining what the key of a value should be.
// Will append to the slice if the key already exists.
func ToSliceMap[T any, K comparable](slice []T, keyFunc func(T) K) map[K][]T {
	m := make(map[K][]T, len(slice))
	for _, v := range slice {
		key := keyFunc(v)
		m[key] = append(m[key], v)
	}
	return m
}
