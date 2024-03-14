package go_basic_radix_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRadixTree(t *testing.T) {
	t.Run("should return a new RadixTree", func(t *testing.T) {
		// ACT
		actual := NewRadixTree()

		var expected *RadixTree = nil
		assert.Equal(t, expected, actual)
	})
}
