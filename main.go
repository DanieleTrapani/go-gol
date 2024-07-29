package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	grid := makeGrid(20, 10)
	for {
		printGrid(grid)
		fmt.Println("\n----------------")
		fmt.Println("")
		grid.UpdateGrid()
		time.Sleep(time.Second)
	}
}

type Grid struct {
	width  int
	height int
	cells  [][]Cell
}

func (g Grid) CountLiveNeighbors(row, col int) int {
	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	count := 0
	for _, dir := range directions {
		r, c := row+dir[0], col+dir[1]
		if r >= 0 && r < len(g.cells) && c >= 0 && c < len(g.cells[0]) && g.cells[r][c].isAlive {
			count++
		}
	}
	return count
}

func (g Grid) UpdateGrid() {
	for row := range g.cells {
		for col := range g.cells[row] {
			// behaviour if cell is alive
			if g.cells[row][col].isAlive {
				aliveNeighbors := g.CountLiveNeighbors(row, col)
				if aliveNeighbors < 2 || aliveNeighbors > 3 {
					g.cells[row][col].isAlive = false
				}
			}
			// behaviour if cell is dead
			if !g.cells[row][col].isAlive && g.CountLiveNeighbors(row, col) == 3 {
				g.cells[row][col].isAlive = true
			}
		}
	}
}

type Cell struct {
	isAlive bool
}

func makeGrid(w, h int) Grid {
	grid := Grid{w, h, make([][]Cell, h)}
	for i := 0; i < h; i++ {
		grid.cells[i] = make([]Cell, w)
	}

	return randomizeGrid(grid)
}

func printGrid(g Grid) {
	for _, row := range g.cells {
		for _, cell := range row {
			if cell.isAlive {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func randomizeGrid(g Grid) Grid {
	for i := range g.cells {
		for j := range g.cells[i] {
			g.cells[i][j].isAlive = randBool()
		}
	}

	return g
}

func randBool() bool {
	num := rand.Float64()
	if num <= 0.15 {
		return true
	} else {
		return false
	}
}
