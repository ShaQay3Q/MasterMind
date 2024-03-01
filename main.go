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
	Indigito
)

type combination [4]color

type hint struct {
	black, white int
}

type entry struct {
	color color
	digit int
}

// compareGuessToAnswer compares the guess and answer and bring backs hints
func compareGuessToAnswer(answer combination, guess combination) hint {

	var res hint
	list := []entry{}

	for i := range answer {
		a := sliceIt(answer, i)
		for j := range guess {
			g := sliceIt(guess, j)
			switch {
			case needBlackHint(a, g, i, j):
				res.black += 1
				break
			case needWhiteHint(a, g, i, j):
				list = addToList(a, j, list)
				res.white = len(addToList(a, j, list))
				break
			}
		}
	}
	return res
}

// addWhiteHint keeps a list of white hints and count the list's length to return
// the number of white hints
func addWhiteHint(color color, i int, list []entry) int {

	list = addToList(color, i, list)
	return len(addToList(color, i, list))
}

// needBlackHint checks for need for black hints
func needBlackHint(color_index_i color, color_index_j color, i int, j int) bool {
	return compare(color_index_i, color_index_j) && i == j
}

// needWhiteHint checks for need for white hints
func needWhiteHint(color_index_i color, color_index_j color, i int, j int) bool {
	return compare(color_index_i, color_index_j) && i != j
}

// compare compares only two slices of the length 1 together
func compare(answer color, guess color) bool {
	return answer == guess
}

// sliceIt get an specific index from combi and parse it into color
func sliceIt(combi combination, i int) color {
	slc := combi[i : i+1]
	color := slc[0]
	return color
}

// addToList keep track of white hints
func addToList(col color, digit int, list []entry) []entry {

	newEntry := entry{col, digit}
	switch {
	case len(list) == 0:
		list = append(list, newEntry)
		return list
	case len(list) != 0:
		if sameColorAndIndex(newEntry, list) {
			break
		} else {
			list = append(list, newEntry)
			break
		}
	}
	return list
}

// sameColorAndIdex check returns true when the colors of certain indexes match
func sameColorAndIndex(newEntry entry, list []entry) bool {
	for i := range list {
		if newEntry.color == list[i].color && newEntry.digit == list[i].digit {
			return true
		}
	}
	return false
}
