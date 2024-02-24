package main

import (
	_ "image/png"

	//"runtime/trace"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompareIndexes01(t *testing.T) {

	answer := combination{Red, Yellow, Blue, Green}
	guess := combination{Red, Yellow, Blue, Green}

	res := compareIndexes(answer, guess)

	require.Equal(t, hint{4, 0}, res)
}

func TestCompareIndexes02(t *testing.T) {

	answer := combination{Red, Yellow, Blue, Green}
	guess := combination{Red, Yellow, Blue, Purple}

	res := compareIndexes(answer, guess)

	require.Equal(t, hint{3, 0}, res)
}

func TestCompareIndexes03(t *testing.T) {

	answer := combination{Red, Yellow, Blue, Green}
	guess := combination{Red, Yellow, Green, Blue}

	res := compareIndexes(answer, guess)

	require.Equal(t, hint{2, 2}, res)
}

func TestCompareIndexes04(t *testing.T) {

	answer := combination{Blue, Yellow, Blue, Green}
	guess := combination{Red, Yellow, Green, Blue}

	res := compareIndexes(answer, guess)

	require.Equal(t, hint{1, 2}, res)
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

func TestAddToList(t *testing.T) {
	var list []entry

	list = addToList(Red, 1, list)

	require.Equal(t, []entry{{Red, 1}}, list)

	list = addToList(Red, 2, list)
	require.Equal(t, []entry{{Red, 1}, {Red, 2}}, list)

	list = addToList(Red, 2, list)
	require.Equal(t, []entry{{Red, 1}, {Red, 2}}, list)

	list = addToList(Red, 1, list)
	require.Equal(t, []entry{{Red, 1}, {Red, 2}}, list)

	list = addToList(Blue, 1, list)
	require.Equal(t, []entry{{Red, 1}, {Red, 2}, {Blue, 1}}, list)
}
