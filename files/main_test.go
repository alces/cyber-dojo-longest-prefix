package common_prefix

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestLongest(t *testing.T) {
    assert.Equal(t, 3, Longest([]string{"newest", "new", "newly"}))
}
