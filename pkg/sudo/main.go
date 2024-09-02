package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 9

// 生成一个完整的数独解
func generateFullSudoku() [N][N]int {
	var grid [N][N]int
	fillDiagonal(&grid)
	fillRemaining(&grid, 0, 3)
	return grid
}

// 填充对角线3x3宫格
func fillDiagonal(grid *[N][N]int) {
	for i := 0; i < N; i += 3 {
		fillBox(grid, i, i)
	}
}

// 填充3x3宫格
func fillBox(grid *[N][N]int, row, col int) {
	nums := rand.Perm(N)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			grid[row+i][col+j] = nums[i*3+j] + 1
		}
	}
}

// 检查数字是否可以放入
func isSafe(grid *[N][N]int, row, col, num int) bool {
	return !usedInRow(grid, row, num) &&
		!usedInCol(grid, col, num) &&
		!usedInBox(grid, row-row%3, col-col%3, num)
}

func usedInRow(grid *[N][N]int, row, num int) bool {
	for col := 0; col < N; col++ {
		if grid[row][col] == num {
			return true
		}
	}
	return false
}

func usedInCol(grid *[N][N]int, col, num int) bool {
	for row := 0; row < N; row++ {
		if grid[row][col] == num {
			return true
		}
	}
	return false
}

func usedInBox(grid *[N][N]int, boxStartRow, boxStartCol, num int) bool {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if grid[row+boxStartRow][col+boxStartCol] == num {
				return true
			}
		}
	}
	return false
}

// 递归填充剩余的格子
func fillRemaining(grid *[N][N]int, i, j int) bool {
	if j >= N && i < N-1 {
		i++
		j = 0
	}
	if i >= N && j >= N {
		return true
	}
	if i < 3 {
		if j < 3 {
			j = 3
		}
	} else if i < N-3 {
		if j == (i/3)*3 {
			j += 3
		}
	} else {
		if j == N-3 {
			i++
			j = 0
			if i >= N {
				return true
			}
		}
	}

	for num := 1; num <= N; num++ {
		if isSafe(grid, i, j, num) {
			grid[i][j] = num
			if fillRemaining(grid, i, j+1) {
				return true
			}
			grid[i][j] = 0
		}
	}
	return false
}

// 递归求解数独，返回解的个数
func countSolutions(grid *[N][N]int, depth int, maxDepth int) int {
	if depth > maxDepth { // 限制递归深度
		return 2 // 如果超过深度，假设有多解
	}

	var i, j int
	found := false
	for i = 0; i < N; i++ {
		for j = 0; j < N; j++ {
			if grid[i][j] == 0 {
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	if !found {
		return 1
	}

	count := 0
	for num := 1; num <= N; num++ {
		if isSafe(grid, i, j, num) {
			grid[i][j] = num
			count += countSolutions(grid, depth+1, maxDepth)
			grid[i][j] = 0
			if count > 1 {
				break // 立即返回，如果发现多解
			}
		}
	}
	return count
}

// 根据难度挖空生成题目，并保证唯一解
func removeDigits(grid *[N][N]int, difficulty float64) {
	numHoles := int(difficulty * float64(N*N)) // 根据难度系数计算要挖掉的格子数量
	count := numHoles
	for count > 0 {
		cellId := rand.Intn(N * N)
		row := cellId / N
		col := cellId % N
		if grid[row][col] != 0 {
			temp := grid[row][col]
			grid[row][col] = 0

			if countSolutions(grid, 0, 20) != 1 { // 假设最大递归深度为20
				grid[row][col] = temp // 恢复原来的值
			} else {
				count--
			}
		}
	}
}

// 打印数独
func printGrid(grid *[N][N]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == 0 {
				fmt.Print("_ ")
			} else {
				fmt.Printf("%d ", grid[i][j])
			}
		}
		fmt.Println()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成数独难度：0.4 表示挖掉40%的格子，适中难度
	difficulty := 0.7

	sudoku := generateFullSudoku()
	removeDigits(&sudoku, difficulty) // 挖掉对应难度系数的数字生成题目
	fmt.Println("生成的数独题目:")
	printGrid(&sudoku)
}
