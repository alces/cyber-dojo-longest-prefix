package common_prefix

func Longest(words []string) int {
    size := len(words)
    
    if size == 0 {
        return 0
    }
    
    firstLen := len(words[0])
    
    for l := 0; l < firstLen; l++ {       
        prefix := getPrefix(words[0], l + 1)

        for i := 1; i < size; i++ {
            if getPrefix(words[i], l + 1) != prefix {
                return l
            }
        }
    }
    
    return firstLen
}

func getPrefix(word string, length int) string {
    if length > len(word) {
        return word
    }
    
    return word[:length]
}