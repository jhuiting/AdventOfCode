package day5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const example = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func Test_part1(t *testing.T) {
	assert.Equal(t, 143, part1(example))
	assert.Equal(t, 5955, part1(input))
}

func Test_part2(t *testing.T) {
	assert.Equal(t, 123, part2(example))
	assert.Equal(t, 4030, part2(input))
}
