// import "slices"
import "strings"

// func dfs(s string, lengths []int, dict map[string]struct{}, words []string, result *[]string) {

// }


func wordBreak(s string, wordDict []string) []string {
    dict := make(map[string]bool, len(wordDict))
    for _, s := range wordDict {
        dict[s] = true
    }
    mem := map[string][]string{}

    var dfs func (substr string) []string
    dfs = func (substr string) []string {
        if len(substr) == 0 {
            return []string{}
        }

        if _, ok := mem[substr]; ok {
            return mem[substr]
        }
        // fmt.Println("substr: ", substr)

        result := []string{}
        if dict[substr] == true {
            result = append(result, substr)
        }

        for i:= 1; i < len(substr); i++ {
            if dict[substr[:i]] == true {
                ss := dfs(substr[i:])
                if len(ss) == 0 {
                    continue
                }
                for _, str := range ss {
                    result = append(result, strings.Join([]string{substr[:i], str}, " "))
                }
            }
        }
        // fmt.Println("result: ",result)
        mem[substr] = result
        return result
    }

    ans := dfs(s)
    fmt.Println(ans)
    return ans
}
