package common_prefix

func Longest(words []string) int {
    size := len(words)
    
    if size == 0 {
        return 0
    }
    
    if size == 1 {
        return len(words[0])
    }
    
    for l := 0; l < len(words[0]); l++ {
        index := l + 1        
        prefix := words[0][:index]

        for i := 1; i < size; i++ {
            if index > len(words[i]) || words[i][:index] != prefix {
                return l
            }
        }
    }
    
    return len(words[0])
}
