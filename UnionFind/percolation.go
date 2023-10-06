package main

type Site int8

const {
    CLOSED Site = 0
    OPEN Site = 1
    FULL Site = 2
}

func CreateGrid(n int) [][]Site {
    var grid [][]Site
    for i := 0; i < n; i++ {
        var row []Site
        for j := 0; j < n; j++ {
            row = append(row, CLOSED)
        }
        grid = append(grid, row)
    }

    return grid
}

func OpenSite(grid *[][]Site, row int, col int) {
    (*grid)[row][col] = OPEN
}

func IsOpen(grid *[][]Site, row int, col int) Site {
    return (*grid)[row][col]
}

func IsFull(grid *[][]Site, row int, col int) bool {
    return (*grid)[row][col] == FULL
}


