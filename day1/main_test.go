package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const example = `3   4
4   3
2   5
1   3
3   9
3   3`

func Test_part1(t *testing.T) {
	assert.Equal(t, 11, part1(example), 11)
	assert.Equal(t, 1873376, part1(input))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 31, part2(example))
	assert.Equal(t, 18997088, part2(input))
}
