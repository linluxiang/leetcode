func uniquePaths(m int, n int) int {
    if m == 1 && n == 1 {
        return 1
    }
  num := make([][]int, m)
  for i:=0; i <m; i++ {
    num[i] = make([]int, n)
  }
  for i:=0; i<m; i++ {
    for j:=0; j <n; j++ {
        if i ==0 && j == 0 {
            num[i][j] = 1
        } else if i == 0 && j != 0 {
            num[i][j] = 1
        } else if i != 0 && j == 0 {
            num[i][j] = 1
        } else {
            num[i][j] = num[i-1][j] + num[i][j-1]
        }
        
    }
  }
  return num[m-1][n-1]
}
