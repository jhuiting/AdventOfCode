package day6

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const example = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func Test_part1(t *testing.T) {
	assert.Equal(t, 41, part1(example))
	assert.Equal(t, 4696, part1(input))
}

func Test_part2(t *testing.T) {
	start := time.Now()
	assert.Equal(t, 6, part2(example))
	durationPart1 := time.Since(start)
	fmt.Println("part 1 duration:", durationPart1)
	assert.Equal(t, 1443, part2(input))
	durationPart2 := time.Since(start)
	fmt.Println("part 2 duration:", durationPart2-durationPart1)

}
