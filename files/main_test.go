package common_prefix

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var longestResults = []struct{
    words    []string
    expected int
} {
    {[]string{"newest", "new", "newly"}, 3},
}

func TestLongest(t *testing.T) {
    for _, r := range longestResults {
        assert.Equal(t, r.expected, Longest(r.words), r.words)
    }
}
