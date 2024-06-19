package common_prefix

func Longest(words []string) (result int) {
    for {
        index := result + 2
        prefix := words[0][:index]
        for i := 1; i < len(words); i++ {
            if words[i][:index] != prefix {
                return
            }
        }
        result += 1
    }
    
    return
}
