import "fmt"

var SET [3]byte = [3]byte{'a', 'b', 'c'}

func generateHappyString(currentString string, cnt *int, length, k int) string {
    // fmt.Println("str: ", currentString, cnt)
    if len(currentString) == length {
        *cnt += 1
        if *cnt == k {
            return currentString
        } else {
            return ""
        }
    }

    for _, letter := range SET {
        if len(currentString) >= 1 && currentString[len(currentString) - 1] == letter {
            continue
        }

        v := generateHappyString(currentString + string(letter), cnt, length, k)
        if v != "" {
            return v
        }

    }
    return ""
}

func getHappyString(n int, k int) string {
    cnt := 0
    return generateHappyString("", &cnt, n, k)
}
