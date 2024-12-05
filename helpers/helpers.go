package helpers

import "strconv"

func StringsToInts(strings []string) ([]int, error) {
	ints := make([]int, len(strings))
	for i, s := range strings {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		ints[i] = num
	}

	return ints, nil
}
