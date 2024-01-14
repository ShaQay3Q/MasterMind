package main

import (
	_ "image/png"

	//"runtime/trace"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompare(t *testing.T) {

	answer := combination{Red, Red, Red, Red}
	guess := combination{Red, Red, Red, Red}

	res := compare(answer, guess)
	require.Equal(t, hint{4, 0}, res)
}
