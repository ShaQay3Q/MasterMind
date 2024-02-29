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

func TestSameColorAndIndex(t *testing.T) {

	list := []entry{{Red, 1}}
	newEntry := entry{Red, 1}

	require.True(t, sameColorAndIndex(newEntry, list))

	list = []entry{{Red, 1}, {Blue, 2}}
	newEntry = entry{Red, 2}

	require.False(t, sameColorAndIndex(newEntry, list))
}

// NOTE: I want to find a way to break down the giveHint into 2 functions. One give back BLACK hints
// the other one return only WHITE hints

func TestNeedBlackHint(t *testing.T) {
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
	i := 1
	j := 1
	a := sliceIt(answer, i)
	g := sliceIt(guess, j)

	require.True(t, needBlackHint(a, g, i, j))

	j = 2
	g = sliceIt(guess, j)
	require.False(t, needBlackHint(a, g, i, j))
}

func TestNeedWhiteHint(t *testing.T) {

	guess := combination{
		Blue,
		Red,
		Yellow,
		Green,
	}

	answer := combination{
		Red,
		Blue,
		Yellow,
		Purple,
	}
	i := 0
	j := 1
	a := sliceIt(answer, i)
	g := sliceIt(guess, j)

	require.True(t, needWhiteHint(a, g, i, j))

	j = 2
	g = sliceIt(guess, j)

	require.False(t, needWhiteHint(a, g, i, j))
}

// func TestAddWhiteHint(t *testing.T) {

// 	i := 1
// 	h := hint{}

// 	h.black = addWhiteHint(i)
// 	require.Equal(t, 2, h.black)
// }
