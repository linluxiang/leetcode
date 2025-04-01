import "strings"

func mergeAlternately(word1 string, word2 string) string {
    i := 0
    builder := strings.Builder{}
    size := min(len(word1), len(word2))
    for i = 0; i < size; i++ {
        builder.WriteByte(word1[i])
        builder.WriteByte(word2[i])
    }

    if i < len(word1) {
        builder.WriteString(word1[i:])
    } else if i < len(word2) {
        builder.WriteString(word2[i:])
    }
    return builder.String()
}
