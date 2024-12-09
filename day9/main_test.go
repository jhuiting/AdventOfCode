package day9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const example = `2333133121414131402`

func Test_part1(t *testing.T) {
	assert.Equal(t, 1928, part1(example))
	assert.Equal(t, 6359213660505, part1(input))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 2858, part2(example))
	assert.Equal(t, 6381624803796, part2(input))
}
