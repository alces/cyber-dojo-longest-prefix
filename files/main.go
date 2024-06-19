package common_prefix

func Longest(words []string) (result int) {
    size := len(words)
    
    if size == 0 {
        return
    }
    
    if size == 1 {
        return len(words[0])
    }
    
    for {
        index := result + 1
        if index > len(words[0]) {
            return
        }
        
        prefix := words[0][:index]
        for i := 1; i < size; i++ {
            if index > len(words[i]) || words[i][:index] != prefix {
                return
            }
        }
        result++
    }
    
    return
}
