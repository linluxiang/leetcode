// import "math"
import "fmt"

// func generateFull(fst int, length int) []int {
//     for p := range length {
//         for n := range math.Pow10(p+1) {
//             fst*math.Pow10(p) + n
//         }
        
//     }
// }

/*
         0
   1       2  3 ...9
 10..19   20..29 ...
 100.109 110..119
 1000..1009


*/

func dfs(n int, num int, result *[]int) {
    if num > n {
        return
    }
    *result = append(*result, num)
    // fmt.Println(*result)
    for v := num*10; v < num*10+10; v++ {
        dfs(n, v, result)
    }
}


func lexicalOrder(n int) []int {
    result := make([]int, 0, n)
    if n < 10 {
        for i:=1; i <=n; i++ {
            result = append(result, i)
        }
        return result
    }

    // fmt.Println(result)
    // for i :=1; i <= 9; i++ {
    //     dfs(n, i, &result)
    // }
    num := 1
    for len(result) <n {
        for num <= n {
            result = append(result, num)
            num *= 10
        }
        // num is too big go back one level
        num /= 10
        for num % 10 == 9 || num > n {
            num /= 10
        }

        num += 1
    }

    return result
}
