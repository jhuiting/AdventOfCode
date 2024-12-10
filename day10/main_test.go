package day10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const exampleTwo = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func Test_part1(t *testing.T) {
	assert.Equal(t, 36, part1(exampleTwo))
	assert.Equal(t, 816, part1(input))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 81, part2(exampleTwo))
	assert.Equal(t, 1960, part2(input))
}
