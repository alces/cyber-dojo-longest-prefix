package common_prefix

func Longest(words []string) int {
    size := len(words)
    
    if size == 0 {
        return 0
    }
    
    firstLen := len(words[0])
    
    for l := 0; l < firstLen; l++ {
        index := l + 1        
        prefix := getPrefix(words[0], index)

        for i := 1; i < size; i++ {
            if getPrefix(words[i], index) != prefix {
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