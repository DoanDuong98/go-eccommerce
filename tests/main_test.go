package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAddOne(t *testing.T) {
	var (
		input  = 1
		output = 3
	)
	//
	//actual := AddOne(input, 2)
	//if actual != output {
	//	t.Errorf("AddOne failed: expected %d, got %d", output, actual)
	//}

	//assert.Equal(t, AddOne(input, 2), output, "AddOne failed")
	assert.Equal(t, AddOne(input, 2), output, "AddOne failed")
	assert.NotEqual(t, AddOne(input, 1), output, "AddOne failed 22")
}

func TestRequire(t *testing.T) {
	require.Equal(t, 2, 3)
	fmt.Print("AAAAAA")
}
