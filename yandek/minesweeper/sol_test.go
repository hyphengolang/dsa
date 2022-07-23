package minesweeper

import "testing"

func TestMineSweeper(t *testing.T) {
	cs := []cooridinate{
		{1, 1},
		{2, 2},
	}

	algo(cs, 3, 2)
}

type cooridinate [2]int

type field map[cooridinate][]cooridinate

func algo(cs []cooridinate, row, col int) {
	f := make(field)

	for _, c := range cs {
		if _, ok := f[c]; !ok {
			f[c] = []cooridinate{}

			// populate the field
			for i := c[0] - 1; i <= c[0]+1; i++ {
				for j := c[1] - 1; j <= c[1]+1; j++ {
					k := cooridinate{i, j}
					if c != k && i > 0 && i <= row && j > 0 && j <= col {
						// f[c] = append(f[c], cooridinate{i, j})
						f[k] = append(f[k], c)
					}
				}

			}
		} else {
			f[c] = []cooridinate{}

			// populate the field
			for i := c[0] - 1; i <= c[0]+1; i++ {
				for j := c[1] - 1; j <= c[1]+1; j++ {
					k := cooridinate{i, j}
					if c != k && i > 0 && i <= row && j > 0 && j <= col {
						// f[c] = append(f[c], cooridinate{i, j})
						f[k] = append(f[k], c)
					}
				}

			}
		}
	}
	_ = f
}
