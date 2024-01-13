package main

import (
	_ "image/png"

	//"runtime/trace"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompare(t *testing.T) {

	var answer [4]color
	var guess [4]color

	answer = [4]color{"R", "R", "R", "R"}
	guess = [4]color{"R", "R", "R", "R"}

	var res hint
	res = compare(answer, guess)
	require.Equal(t, 4, res.b)
}
