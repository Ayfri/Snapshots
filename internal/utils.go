package internal

func Any[T comparable](array []T, predicate func(T) bool) bool {
	for _, v := range array {
		if predicate(v) {
			return true
		}
	}

	return false
}

func Contains[T comparable](array []T, value T) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}

	return false
}

func Filter[T comparable](array []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range array {
		if predicate(v) {
			result = append(result, v)
		}
	}

	return result
}

func Find[T comparable](array []T, predicate func(T) bool) (T, bool) {
	for _, v := range array {
		if predicate(v) {
			return v, true
		}
	}

	var result T
	return result, false
}

func GroupBy[T comparable, V comparable](array []T, transformable func(T) V) map[V][]T {
	result := make(map[V][]T)
	for _, v := range array {
		key := transformable(v)
		result[key] = append(result[key], v)
	}

	return result
}

func None[T comparable](array []T, predicate func(T) bool) bool {
	for _, v := range array {
		if predicate(v) {
			return false
		}
	}

	return true
}
