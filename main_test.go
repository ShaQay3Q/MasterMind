package main

import (
	_ "image/png"

	//"runtime/trace"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompare(t *testing.T) {

	answer := combination{"R", "R", "R", "R"}
	guess := combination{"R", "R", "R", "R"}

	res := compare(answer, guess)
	require.Equal(t, hint{4, 0}, res)
}
