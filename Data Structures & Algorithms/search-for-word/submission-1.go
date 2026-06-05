func exist(board [][]byte, word string) bool {
    rows, cols := len(board), len(board[0])
    visited := make([][]bool, rows)
    for i := range visited {
        visited[i] = make([]bool, cols)
    }

    var dfs func(r, c, i int) bool
    dfs = func(r, c, i int) bool {
        if i == len(word) {
            return true
        }

        if r < 0 || c < 0 || r >= rows || c >= cols ||
            board[r][c] != word[i] || visited[r][c] {
            return false
        }

        visited[r][c] = true
        res := dfs(r+1, c, i+1) ||
                dfs(r, c+1, i+1) ||
                dfs(r-1, c, i+1) ||
                dfs(r, c-1, i+1)
        visited[r][c] = false

        return res
    }

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if dfs(r, c, 0) {
                return true
            }
        }
    }

    return false
}


























// hashset
    // rows := len(board) 
    // cols := len(board[0])
    // visited := make(map[[2]int]bool)
    // var dfs func(r, c, i int) bool
    // dfs = func(r, c, i int) bool {
    //     if (i == len(word)) {
    //         return true
    //     }

    //     if r < 0 || r >= rows || c < 0 || c >= cols || 
    //     visited[[2]int{r,c}] ||
    //     board[r][c] != word[i] {
    //         return false
    //     }

    //     visited[[2]int{r, c}] = true
    //     res := dfs(r + 1, c, i + 1) ||
    //             dfs(r, c + 1, i + 1) ||
    //             dfs(r - 1, c, i + 1) ||
    //             dfs(r, c - 1, i + 1)
    //     visited[[2]int{r, c}] = false

    //     return res
    // }

    // for i := 0; i < rows; i++ {
    //     for j := 0; j < cols; j++ {
    //         if dfs(i, j, 0) {
    //             return true
    //         }
    //     }
    // }

    // return false