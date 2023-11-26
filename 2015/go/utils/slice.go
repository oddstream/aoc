package utils

func Contains[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func Union[T comparable](a []T, b []T) []T {
	var result []T = a
	for _, elem := range b {
		if !Contains(a, elem) {
			result = append(result, elem)
		}
	}
	return result
}

// brute force (inefficient) method O(m∗n)
func IntersectionSlow[T comparable](a, b []T) []T {
	var result []T
	for _, elem := range a {
		if Contains(b, elem) {
			result = append(result, elem)
		}
	}
	return result
}

// more efficient method using a set O(m+n)
func Intersection[T comparable](a, b []T) []T {
	// create an empty result array, with enough room to hold intersection
	var result = make([]T, 0, len(a))
	// create a set from the first array
	var set map[T]struct{} = make(map[T]struct{})
	for _, ele := range a {
		set[ele] = struct{}{}
	}
	// check if each element exists in the set by traversing through the second array
	for _, ele := range b {
		if _, ok := set[ele]; ok {
			// if an element exists in both arrays, add it to the intersection array
			result = append(result, ele)
		}
	}
	return result
}

func Exclusion[T comparable](a []T, b []T) []T {
	var result []T
	for _, elem := range a {
		if !Contains(b, elem) {
			result = append(result, elem)
		}
	}
	return result
}
