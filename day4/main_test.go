package day4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const example = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func Test_part1(t *testing.T) {
	assert.Equal(t, 18, part1(example))
	assert.Equal(t, 2662, part1(input))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 9, part2(example))
	assert.Equal(t, 2034, part2(input))
}
