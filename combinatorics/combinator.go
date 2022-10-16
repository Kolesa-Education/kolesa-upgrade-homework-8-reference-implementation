package combinatorics

func deduplicate[T comparable](elements []T) []T {
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
