package utils

import "slices"

func MachesAny[T comparable](as, bs []T) bool {
	return slices.ContainsFunc(as,
		func(a T) bool {
			return slices.Contains(bs, a)
		})
}

func Map[X, Y any](xs []X, foo func(X) Y) []Y {
	ys := make([]Y, 0, len(xs))

	for _, x := range xs {
		ys = append(ys, foo(x))
	}

	return ys
}

// DSNTE Dereference string nil to empty
func DSNTE(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
