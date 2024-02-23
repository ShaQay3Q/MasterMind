package main

func main() {

}

type color int

const (
	Red color = iota
	Blue
	Green
	Yellow
	Purple
	Indigo
)

type combination [4]color

type hint struct {
	black, white int
}

func pop(combination combination, i int) [1]color {
	var res [1]color
	res[i] = combination[i]
	return res
}

// func compareGuessToAnswer(answer combination, guess combination) hint {

// 	res := compareSameIndexes(answer, guess)

// 	return res
// }

func compareSameIndexes(answer combination, guess combination) hint {
	var res hint

	for i := range answer {
		a := sliceIt(answer, i)
		g := sliceIt(guess, i)
		if compare(a, g) {
			res.black += 1
		}
	}

	return res
}

// compare compares only two slices of the length 1 together
func compare(answer color, guess color) bool {

	if answer == guess {
		return true
	}
	return false
}

// sliceIt get an specific index from combi and parse it into color
func sliceIt(combi combination, i int) color {
	slc := combi[i : i+1]
	color := slc[0]
	return color
}
