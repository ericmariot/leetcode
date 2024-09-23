func orangesRotting(grid [][]int) int {
    rotten := [][]int{}
    size := 0
    rottenSize := 0
    res := 0

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 1 {
                size++
            }
            if grid[i][j] == 2 {
                rotten = append(rotten, []int{i, j})
                size++
                rottenSize++
            }
        }
    }

    for len(rotten) > 0 {
        if size == rottenSize {
            break
        }

        l := len(rotten)

        for range l {
            o := rotten[0]
            rotten = rotten[1:]
            if o[0] < len(grid)-1 {
                i, j := o[0]+1, o[1]
                if grid[i][j] == 1 {
                    grid[i][j] = 2
                    rotten = append(rotten, []int{i, j})
                    rottenSize++
                }
            }

            if o[0] > 0 {
                i, j := o[0]-1, o[1]
                if grid[i][j] == 1 {
                    grid[i][j] = 2
                    rotten = append(rotten, []int{i, j})
                    rottenSize++
                }
            }

            if o[1] > 0 {
                i, j := o[0], o[1]-1
                if grid[i][j] == 1 {
                    grid[i][j] = 2
                    rotten = append(rotten, []int{i, j})
                    rottenSize++
                }
            }

            if o[1] < len(grid[0])-1 {
                i, j := o[0], o[1]+1
                if grid[i][j] == 1 {
                    grid[i][j] = 2
                    rotten = append(rotten, []int{i, j})
                    rottenSize++
                }
            }
        }
        
        res++
    }

    if rottenSize != size {
        return -1
    }

    return res
}
