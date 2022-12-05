package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func TestForWindows(t *testing.T) {
	res, err := WifiData(runtime.GOOS)
	assert.NoError(t, err)

	fmt.Printf("strength %s \n", res.PrettyPrint())
}
