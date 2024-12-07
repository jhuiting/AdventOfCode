package day7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const example = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func Test_part1(t *testing.T) {
	assert.Equal(t, 3749, part1(example))
	assert.Equal(t, 42283209483350, part1(input))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 11387, part2(example))
	assert.Equal(t, 1026766857276279, part2(input))
}
