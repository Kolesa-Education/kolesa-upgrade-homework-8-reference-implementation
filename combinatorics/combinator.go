package combinatorics

import (
	combo "github.com/natemcintosh/gocombinatorics"
)

func Deduplicate[T comparable](elements []T) []T {
	allKeys := map[T]bool{}
	var list []T

	for _, item := range elements {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func Combinations[T comparable](elements []T, sliceSize int) ([][]T, error) {
	combination, err := combo.NewCombinations(elements, sliceSize)
	if err != nil {
		return nil, err
	}
	var result [][]T
	for combination.Next() {
		items := combination.Items()
		copied := make([]T, len(items))
		copy(copied, items)
		result = append(result, copied)
	}
	return result, nil
}
