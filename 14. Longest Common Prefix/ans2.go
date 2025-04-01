func longestCommonPrefix(strs []string) string {
    if len(strs) == 1 {
        return strs[0]
    }
    for i, c := range []byte(strs[0]) {
        for j :=1; j < len(strs); j++ {
            s := strs[j]
            if i == len(s) || c != s[i] {
                return strs[0][:i]
            }
        }
    }
    return strs[0]
}
