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

func Reverse[E any](s []E) []E {
	result := make([]E, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}
	return result
}

/*
func Map[E any](s []E, f mapFunc[E]) []E {
    result := make([]E, len(s))
    for i := range s {
        result[i] = f(s[i])
    }
    return result
}

func Filter[E any](s []E, f keepFunc[E]) []E {
    result := []E{}
    for _, v := range s {
        if f(v) {
            result = append(result, v)
        }
    }
    return result
}

func Reduce[E any](s []E, init E, f reduceFunc[E]) E {
    cur := init
    for _, v := range s {
        cur = f(cur, v)
    }
    return cur
}
*/
