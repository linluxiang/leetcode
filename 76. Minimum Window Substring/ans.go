
// 最粗略版本，比较字典中是否相同
func equal(appears, window map[byte]int) bool {
    if len(appears) != len(window) {
        return false
    }

    for k := range appears {
        if window[k] < appears[k] {
            return false
        }
    }

    return true
}

func minWindow(s string, t string) string {
    appears := map[byte]int{}
    for _, c := range []byte(t) {
        appears[c] += 1
    }

    window := map[byte]int{}
    valid := 0 // valid 用来记录当前window中，有多少个数量和t中一样的字符
    minlen := len(s) + 1
    i, j := 0, 0
    var mini, minj int
    for j < len(s) {
        char := s[j]
        j++

        if appears[char] != 0 {
            window[char] += 1
            if window[char] == appears[char] { // 关键点，如果当前字符满足要求，就valid + 1，这样只需要记录多少个满足要求的字符就可以了。
                valid += 1
            }
        }

        // fmt.Println("window: [", i, ", ", j, ")", window, s[i:j])
        // for i < j && equal(appears, window) {
        for i < j && valid == len(appears) {
            if j - i < minlen {
                minlen = j - i
                mini, minj = i, j
            }

            c := s[i]
            if appears[c] != 0 {
                if window[c] == appears[c] {
                    valid -= 1
                }
                window[c] -= 1

            }
            i++
            // fmt.Println("narrow left : [", i, ", ", j, ")", window, s[i:j])
        }
    }

    if minlen == len(s) + 1 {
        return ""
    }

    return s[mini:minj]
} 
