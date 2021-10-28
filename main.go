package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

const (
	Width  = 80
	Height = 30
)

func main() {

	var grid [Height][Width]int = [Height][Width]int{}
	var buffer [Height][Width]int = [Height][Width]int{}

	//generating random starting sequence
	for x := 1; x < Height-1; x++ {
		for y := 1; y < Width-1; y++ {
			if rand.Float32() < 0.5 {
				grid[x][y] = 1
			}
		}
	}
	tru := true
	for tru {
		fmt.Printf("\n=======NEW GENERATION =======")

		//printing
		for i := 1; i < Height-1; i++ {
			fmt.Printf("\n")
			for j := 1; j < Width-1; j++ {
				if grid[i][j] == 1 {
					fmt.Printf("⬜")
				} else {
					fmt.Printf("⬛")
				}

			}
		}
		// killing/reproducing cells // i = x, j = y
		for i := 1; i < Height-1; i++ {
			for j := 1; j < Width-1; j++ {
				buffer[i][j] = 0
				//finding neighbors
				n := grid[i-1][j-1] + grid[i-1][j+0] + grid[i-1][j+1] + grid[i+0][j-1] + grid[i+0][j+1] + grid[i+1][j-1] + grid[i+1][j+0] + grid[i+1][j+1]
				if grid[i][j] == 0 && n == 3 {
					buffer[i][j] = 1
				} else if n > 3 || n < 2 {
					buffer[i][j] = 0
				} else {
					buffer[i][j] = grid[i][j]
				}
			}
		}
		temp := buffer
		buffer = grid
		grid = temp

		fmt.Printf("\nPRESS ENTER TO CONTINUE")

		//getchar
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		if input == "100" {
			fmt.Println(input)
		}

	}

}
