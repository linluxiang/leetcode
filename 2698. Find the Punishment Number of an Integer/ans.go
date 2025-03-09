import "slices"
import "fmt"

func length(target int) int {
    c := 0
    if target < 10 {
        return 1
    }
    for target != 0 {
        target = target / 10
        c += 1
    }
    return c
}

func calNum(letters []int) int{
    result := 0
    for i := 0; i < len(letters); i++ {
        result = result * 10 + letters[i]
    }
    return result
}

func cal(letters []int, target int) bool {
    if len(letters) == 0 && target == 0 {
        return true
    }
    // fmt.Println(letters, target)

    targetLength := length(target)
    // fmt.Println("target len:", targetLength)
    for i := targetLength; i >= 1; i-- {
        if i > len(letters) {
            continue
        }
        prefix := calNum(letters[:i])
        // fmt.Println("prefix:", prefix, letters[:i])
        v := cal(letters[i:], target - prefix)
        if v {
            return true
        }

    }
    return false
}

func allLetters(n int) []int {
    nums := []int{}
    x := n
    for x != 0 {
        nums = append(nums, x % 10)
        x = x / 10
    }
    slices.Reverse(nums)
    return nums
}

func isPunishNumber(n int) bool {
    square := n * n
    al := allLetters(square)
    return cal(al, n)
}


func punishmentNumber(n int) int {
    result := 0
    for i := 1; i <= n ; i++ {
        if isPunishNumber(i) {
            // fmt.Println("get: ", i)
            result += i * i
        }
    }

    return result
}
