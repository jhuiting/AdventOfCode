package day11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const example = `125 17`

func Test_part1(t *testing.T) {
	assert.Equal(t, 55312, part1(example))
	assert.Equal(t, 216996, part1(input))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 65601038650482, part2(example))
	assert.Equal(t, 257335372288947, part2(input))
}
