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
        []string{"a", "b"},
        [][]string{{"a", "b"}},
    },
}

func TestAllPairs(t *testing.T) {
    for _, r := range allPairsResults {
        assert.ElementsMatch(t, r.expected, allPairs(r.argument), r.argument)
    }
}