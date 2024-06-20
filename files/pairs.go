package common_prefix

func LongestInPairs(words []string) (result int) {
    for _, p := range allPairs(words) {
        candidate := Longest(p)
        if candidate > result {
            result = candidate
        }
    }
    
    return
}

func allPairs(words []string) (result [][]string) {
    size := len(words)
    
    if size < 2 {
        return
    }
    
    for i := 0; i < size; i++ {
        for j := i + 1; j < size; j++ {
            result = append(result, []string{words[i], words[j]})
        }    
    }    
    
    return
}