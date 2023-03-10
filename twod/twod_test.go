package twod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMove(t *testing.T) {
	assert := assert.New(t)
	type test struct {
		loc      Location
		dir      Direction
		expected Location
	}
	tests := []test{
		{
			loc:      Location{1, 1},
			dir:      Direction{-1, 1},
			expected: Location{0, 2},
		},
		{
			loc:      Location{3, 4},
			dir:      Direction{1, -1},
			expected: Location{4, 3},
		},
		{
			loc:      Location{-1, -2},
			dir:      Direction{4, 5},
			expected: Location{3, 3},
		},
	}
	predefinedDirs := []Direction{
		LEFT, RIGHT, UP, DOWN, N, S, E, W, NE, NW, SE, SW,
	}
	for _, predefinedDir := range predefinedDirs {
		tests = append(tests, test{
			loc:      ORIGIN,
			dir:      predefinedDir,
			expected: Location{predefinedDir[0], predefinedDir[1]},
		})
	}
	for _, t := range tests {
		assert.Equal(t.expected, Move(t.loc, t.dir))
	}
}

func testAddDirs(t *testing.T) {
	assert := assert.New(t)
	type test struct {
		dir1     Direction
		dir2     Direction
		expected Direction
	}
	tests := []test{
		{
			dir1:     Direction{1, 1},
			dir2:     Direction{-1, 1},
			expected: Direction{0, 2},
		},
		{
			dir1:     Direction{3, 4},
			dir2:     Direction{1, -1},
			expected: Direction{4, -3},
		},
		{
			dir1:     Direction{-1, -2},
			dir2:     Direction{4, 5},
			expected: Direction{3, 3},
		},
	}

	for _, t := range tests {
		assert.Equal(t.expected, AddDirs(t.dir1, t.dir2))
	}
}
