package main

import (
	_ "image/png"

	//"runtime/trace"
	"testing"

	"github.com/stretchr/testify/require"
)

// func TestCompareGuessToAnswer01(t *testing.T) {

// 	answer := combination{Red, Red, Red, Red}
// 	guess := combination{Red, Red, Red, Red}

// 	res := compare(answer, guess)

// 	require.Equal(t, hint{4, 0}, res)
// }

func TestCompareSameIndexes(t *testing.T) {

	answer := combination{Red, Yellow, Blue, Green}
	guess := combination{Red, Yellow, Blue, Purple}

	res := compareSameIndexes(answer, guess)

	require.Equal(t, hint{4, 0}, res)
}

// func estCompareGuessToAnswer02(t *testing.T) {

// 	answer := combination{Red, Yellow, Blue, Green}
// 	guess := combination{Red, Yellow, Blue, Purple}

// 	res := compare(answer, guess)

// 	require.Equal(t, hint{3, 0}, res)
// }

func TestCompareSameIndexes02(t *testing.T) {

	answer := combination{Red, Yellow, Blue, Green}
	guess := combination{Red, Yellow, Blue, Purple}

	res := compareSameIndexes(answer, guess)

	require.Equal(t, hint{3, 0}, res)
}

func TestPop(t *testing.T) {
	combi := combination{Red, Yellow, Blue, Green}

	res := combi[1]

	require.Equal(t, Yellow, res)
}

func TestCompare(t *testing.T) {
	guess := combination{
		Red,
		Blue,
		Yellow,
		Green,
	}

	answer := combination{
		Red,
		Blue,
		Yellow,
		Green,
	}
	var a = sliceIt(guess, 1)
	var g = sliceIt(answer, 1)

	require.True(t, compare(a, g))
}

func TestSliceIt(t *testing.T) {
	//var slc []color
	var i int
	combi := combination{
		Red,
		Blue,
		Yellow,
		Green,
	}
	slc := sliceIt(combi, i)

	require.Equal(t, Red, slc)
}
