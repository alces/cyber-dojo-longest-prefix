package common_prefix

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestAllPairs(t *testing.T) {
    argument := []string{"a", "b"}
    expected := [][]string{{"a", "b"}}
    
    assert.ElementsMatch(t, expected, allPairs(argument), argument)
}