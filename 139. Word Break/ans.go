import "fmt"

func search(s string, dict map[string]bool, visited map[string]bool) bool {
    //fmt.Println(s)
    if _, ok := visited[s]; ok {
        return visited[s]
    }

    if dict[s] == true {
        visited[s] = true
        return true
    }

    //fmt.Println("search: ", s)
    for i := len(s)-1; i >=0; i-- {
        if ok := dict[s[:i+1]]; ok {
            //fmt.Println("found:", s[:i+1])
            if search(s[i+1:len(s)], dict, visited) {
                visited[s[i+1:len(s)]] = true
                return true
            }
        }
    }
    visited[s] = false
    return false
}

func wordBreak(s string, wordDict []string) bool {
    dict := map[string]bool {}
    visited := map[string]bool {}
    letters := map[rune]bool {}

    for _, str := range wordDict{
        dict[str] = true
        for _, l := range str {
            letters[l] = true
        }
    }

    for _, l := range s {
        if letters[l] == false{
            return false
        }
    }
    return search(s, dict, visited)
}
