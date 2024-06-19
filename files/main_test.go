package common_prefix

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var longestResults = []struct{
    words    []string
    expected int
} {
    {[]string{}, 0},
    {[]string{"oldest"}, 6},
    {[]string{"newest", "new", "newly"}, 3},
    {[]string{"new", "next"}, 2},
    {[]string{"pond", "pod", "new", "newest"}, 0},
}

func TestLongest(t *testing.T) {
    for _, r := range longestResults {
        assert.Equal(t, r.expected, Longest(r.words), r.words)
    }
}

var getPrefixResults = []struct{
    word     string
    length   int
    expected string
} {
    {"pond", 2, "po"},
}

func TestGetPrefix(t *testing.T) {
    for _, r := range getPrefixResults {
        assert.Equal(t, r.expected, getPrefix(r.word, r.length))
    }
}