package day

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const example = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
const example2 = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func Test_part1(t *testing.T) {
	assert.Equal(t, 161, part1(example))
	assert.Equal(t, 190604937, part1(input))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 48, part2(example2))
	assert.Equal(t, 82857512, part2(input))
}
