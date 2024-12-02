package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const example = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func Test_part1(t *testing.T) {
	assert.Equal(t, 2, part1(example))
	assert.Equal(t, 549, part1(input))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 4, part2(example))
	assert.Equal(t, 589, part2(input))
}
