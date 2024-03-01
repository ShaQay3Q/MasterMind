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

// something broken in here.
// my assumption is that the switch-case part s not working as I expected
// i assume the code was broken and I didn't realize it earlier
// NEED TO DEBUG THIS BIT LATER

func compareIndexes(answer combination, guess combination) hint {

	var res hint
	list := []entry{}

	for i := range answer {
		a := sliceIt(answer, i)
		for j := range guess {
			g := sliceIt(guess, j)
			switch {
			case needBlackHint(a, g, i, j):
				addBlackHint(res.black)
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

func addBlackHint(i int) int {
	return i + 1
}

func addWhiteHint(color color, i int, list []entry) int {

	list = addToList(color, i, list)
	return len(addToList(color, i, list))
}

func needBlackHint(color_index_i color, color_index_j color, i int, j int) bool {
	return compare(color_index_i, color_index_j) && i == j
}

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

func sameColorAndIndex(newEntry entry, list []entry) bool {
	for i := range list {
		if newEntry.color == list[i].color && newEntry.digit == list[i].digit {
			return true
		}
	}
	return false
}

// NOTE: I want to find a way to break down the giveHint into 2 functions. One give back BLACK hints
// the other one return only WHITE hints

// func giveHint(indexColor color, indexNumber int, guess combination) hint {
// 	return hint{4, 0}
// }
