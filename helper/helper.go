package helper

import (
	"golang.org/x/exp/constraints"
)

func Must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}

	return x
}

func Map[T any, R any](input []T, fn func(T) R) []R {
	result := make([]R, 0, len(input))

	for _, v := range input {
		result = append(result, fn(v))
	}

	return result
}

func InGroupsOf[T any](input []T, n int) [][]T {
	result := make([][]T, 0)

	for i := 0; i < len(input); i += n {
		result = append(result, input[i:i+n])
	}

	return result
}

func First[T any](input []T) T {
	return input[0]
}

func Last[T any](input []T) T {
	return input[len(input)-1]
}

func Min[T constraints.Ordered](input []T) T {
	m := input[0]

	for _, v := range input[1:] {
		if v < m {
			m = v
		}
	}

	return m
}

func Max[T constraints.Ordered](input []T) T {
	m := input[0]

	for _, v := range input[1:] {
		if v > m {
			m = v
		}
	}

	return m
}

func All[T any](input []T, fn func(T) bool) bool {
	for _, v := range input {
		if !fn(v) {
			return false
		}
	}

	return true
}

func Any[T any](input []T, fn func(T) bool) bool {
	for _, v := range input {
		if fn(v) {
			return true
		}
	}

	return false
}

func Reduce[T any](input []T, fn func(T, T) T) T {
	result := input[0]

	for _, v := range input[1:] {
		result = fn(result, v)
	}

	return result
}

func Reverse[T any](input []T) []T {
	result := make([]T, len(input))

	for i, v := range input {
		result[len(input)-i-1] = v
	}

	return result
}

func Contains[T comparable](input []T, v T) bool {
	for _, x := range input {
		if x == v {
			return true
		}
	}

	return false
}
