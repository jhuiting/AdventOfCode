package helpers

import "strconv"

func ToInts(strings []string) []int {
	ints := make([]int, len(strings))
	for i, s := range strings {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		ints[i] = num
	}

	return ints
}
