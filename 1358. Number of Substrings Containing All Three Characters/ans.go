func numberOfSubstrings(s string) int {
    existed := [3]int{}
    for _, b := range []byte(s) {
        existed[b-'a'] += 1
    }

    if existed[0] == 0 || existed[1] == 0 || existed[2] == 0 {
        return 0
    }

    sum := 0
    appears := [3]int{0, 0, 0}
    //这其实是个滑动窗口，右指针向右滑动知道同时出现abc，然后左指针向右滑动直到abc缺一个，期间从i到字符串尾巴的所有字符串都可以认为符合要求
    i := 0
    j := -1
    appears[0], appears[1], appears[2] = 0, 0, 0

    // appears[s[0]-'a'] = 1
    for j< len(s) - 1 {
        // 窗口向右扩展，直到找到abc都出现的时候
        for j < len(s) - 1 {
            j++
            appears[s[j] - 'a'] += 1
            if appears[0] != 0 && appears[1] != 0 && appears[2] != 0 {
                // fmt.Println("expand right to :", i, j, appears)
                break
            }
        }

        if appears[0] != 0 && appears[1] != 0 && appears[2] != 0 {
            sum += len(s) - j
        }
        // fmt.Println("sum: ", sum)

        // 窗口从左缩小，直到i和j重合
        for i < j {
            appears[s[i]-'a'] -= 1
            i++
            if appears[0] == 0 || appears[1] == 0 || appears[2] == 0 {
                // fmt.Println("narrow left to :", i, j, appears)
                break
            }
            sum += len(s) - j
            // fmt.Println("next sum: ", sum, i, appears)

        }

    }
    return sum
}
