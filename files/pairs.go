package common_prefix

func allPairs(words []string) (result [][]string) {
    size := len(words)
    
    if size < 2 {
        return
    }
    
    for i := 0; i < size; i++ {
        for j := 0; j < size; j++ {
            if i != j {
                result = append(result, []string{words[i], words[j]})
            }
        }    
    }    
    
    return
}