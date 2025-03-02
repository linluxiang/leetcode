import "fmt"

func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    }

    if len(s) == 1 {
        return 1
    }
    set := map[byte]bool{}
    max := 1
    i, j := 0, 0
    for i < len(s) && j < len(s) {
        // fmt.Println(i, j, set, max)
        if _, ok := set[s[j]]; !ok {
            set[s[j]] = true

            length := j - i + 1
            if length > max {
                max = length
                // fmt.Println("set max: ", max, i, j)
            }
            
            j++
            continue
        }

        // found repeated letter

        for s[i] != s[j] {
            delete(set, s[i])
            i++
        }
        // i is at the index of the repeated letter

        delete(set, s[j])
        i++
    }

    return max
}
