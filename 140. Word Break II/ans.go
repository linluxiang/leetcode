import "slices"
import "strings"

func dfs(s string, lengths []int, dict map[string]struct{}, words []string, result *[]string) {
    // fmt.Println(s, lengths, words)
    for _, length := range lengths {
        if len(s) < length {
            break
        }
        if len(s) == length {
            if _, ok := dict[s]; ok {
                ans := strings.Join(append(words, s), " ")
                *result = append(*result, ans)
                continue
            }
        }

        prefix := s[:length]
        left := s[length:]
        if _, ok := dict[prefix]; ok {
            dfs(left, lengths, dict, append(words, prefix), result)
        }
    }
}


func wordBreak(s string, wordDict []string) []string {
    dict := map[string]struct{}{}
    lengths := make([]int, len(wordDict)) // the lengths of the word in dict, sorted from small to big

    for i, word := range wordDict {
        dict[word] = struct{}{}
        lengths[i] = len(word)
    }
    slices.Sort(lengths)
    lengths = slices.Compact(lengths)


    result := []string{}
    dfs(s, lengths, dict, []string{}, &result)
    return result
}
