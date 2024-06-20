package common_prefix

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

var allPairsResults = []struct {
    argument []string
    expected [][]string
} {
    {
        []string{"a"},
        [][]string{},
    },
    {
        []string{"a", "b", "c"},
        [][]string{{"a", "b"}, {"a", "c"}, {"b", "c"}},
    },
}

func TestAllPairs(t *testing.T) {
    for _, r := range allPairsResults {
        assert.ElementsMatch(t, r.expected, allPairs(r.argument), r.argument)
    }
}

var longestInPairsResults = []struct {
    argument []string
    expected int
} {
    {[]string{"pond", "pod", "new", "newest"}, 3},
}

func TestLongestInPairs(t *testing.T) {
    for _, r := range longestInPairsResults {
        assert.Equal(t, r.expected, LongestInPairs(r.argument), r.argument)
    }
}