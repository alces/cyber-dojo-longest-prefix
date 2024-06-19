package common_prefix

func Longest(words []string) (result int) {
    for {
        index := result + 1
        if index > len(words[0]) {
            return
        }
        
        prefix := words[0][:index]
        for i := 1; i < len(words); i++ {
            if index > len(words[i]) || words[i][:index] != prefix {
                return
            }
        }
        result += 1
    }
    
    return
}
