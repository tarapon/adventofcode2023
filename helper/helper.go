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
