package main

type Node struct {
	value int
	left  *Node
	right *Node
	up    *Node
	down  *Node
}

func main() {

}

func maximumMinutes(grid [][]int) int {
    m, n:= len(grid), len(grid[0])
	for i := 0 ; i < m ; i ++ {
		for j := 0, j < n; i ++ {
			node := &Node{
				value: grid[i][j]
			}
		}
	}
}