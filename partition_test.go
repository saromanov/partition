package partition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	value := []int{5, 2, 8, 9, 5, 3, 2, 0, 0, 2, 5, 5}
	for idx := range Partition(len(value), 2) {
		assert.Equal(t, len(value[idx.Start:idx.End]), 2, "not equal")
	}
}
