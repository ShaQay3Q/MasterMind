package main

import (
	_ "image/png"

	//"runtime/trace"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompare01(t *testing.T) {

	answer := combination{Red, Red, Red, Red}
	guess := combination{Red, Red, Red, Red}

	res := compare(answer, guess)

	require.Equal(t, hint{4, 0}, res)
}

func TestCompareSameIndexes(t *testing.T) {

	answer := combination{Red, Yellow, Blue, Green}
	guess := combination{Red, Yellow, Blue, Purple}

	res := compareSameIndexes(answer, guess)

	require.Equal(t, hint{4, 0}, res)
}

func TestCompare02(t *testing.T) {

	answer := combination{Red, Yellow, Blue, Green}
	guess := combination{Red, Yellow, Blue, Purple}

	res := compare(answer, guess)

	require.Equal(t, hint{3, 0}, res)
}

func TestCompareSameIndexes02(t *testing.T) {

	answer := combination{Red, Yellow, Blue, Green}
	guess := combination{Red, Yellow, Blue, Purple}

	res := compareSameIndexes(answer, guess)

	require.Equal(t, hint{3, 0}, res)
}
