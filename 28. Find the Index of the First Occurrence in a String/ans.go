// func buildLps(pattern string) []int {
//     lps := make([]int, len(pattern))
//     lps[0] = 0
//     matchedPrefix := 0
//     i := 1
//     for i < len(pattern) {
//         if pattern[i] == pattern[matchedPrefix] {
//             matchedPrefix += 1
//             lps[i] = matchedPrefix
//             i += 1
//         } else {
//             if matchedPrefix != 0 {
//                 matchedPrefix = lps[matchedPrefix - 1]
//             } else {
//                 lps[i] = 0
//                 i += 1
//             }
//         }
//     }
//     return lps
// }

import "hash/maphash"

func RabinKarp(haystack, needle string) int {
    var h maphash.Hash
    h.WriteString(needle)
    hashPattern := h.Sum64()

    for i := range len(haystack) {
        if i+len(needle) > len(haystack) {
            return -1
        }
        h.Reset()
        h.WriteString(haystack[i:i+len(needle)])
        hv := h.Sum64()
        if hashPattern == hv {
            if needle == haystack[i:i+len(needle)] {
                return i
            }
        }
    }
    return -1
}

func strStr(haystack string, needle string) int {
    if len(needle) > len(haystack) {
        return -1
    }
    return RabinKarp(haystack, needle)
}
