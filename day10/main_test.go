package day10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const example = `0123
1234
8765
9876`

const exampleTwo = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func Test_part1(t *testing.T) {
	assert.Equal(t, 0, part1(example))
	// 9 trails
	assert.Equal(t, 36, part1(exampleTwo))
	//assert.Equal(t, 0, part1(input))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 1, part2(example))
}
