package main

import "fmt"

func cellDrawer(cellStatus bool) string {
	if cellStatus {
		return "0"
	} else {
		return "."
	}
}

func main() {
	col := 8
	row := 8
	iter := 5

	cell := make([][]bool, col)
	for i := range col {
		cell[i] = make([]bool, row)
	}

	prem := [][]int{{2, 2}, {3, 2}, {3, 3}, {2, 3}, {1, 3}}

	for i := range prem {
		px := prem[i][0]
		py := prem[i][1]

		cell[px][py] = true
	}

	for n := range iter {
		fmt.Printf("Iteration: %d\n", n)
		for i := range col {
			for j := range row {
				fmt.Printf("%s ", cellDrawer(cell[i][j]))
			}
			fmt.Println()
		}
	}

}
