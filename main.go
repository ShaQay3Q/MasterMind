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

type entry struct {
	color color
	dig   int
}

// func compareGuessToAnswer(answer combination, guess combination) hint {

// 	res := compareSameIndexes(answer, guess)

// 	return res
// }

func compareSameIndexes(answer combination, guess combination) hint {
	var res hint

	for i := range answer {
		a := sliceIt(answer, i)
		for j := range guess {
			g := sliceIt(guess, j)
			if compare(a, g) && i == j {
				res.black += 1
				break
			} else if compare(a, g) && i != j {
				list := []entry{}
				res.white = len(addToList(a, j, list))
				break
			}
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

func addToList(col color, dig int, input []entry) []entry {
	list := input
	newEntry := entry{col, dig}
	if len(list) == 0 {
		list = append(list, newEntry)
	} else {
		for i := 0; i < len(list); i++ {
			if col == list[i].color && dig == list[i].dig {
				return list
			}
		}
		list = append(list, newEntry)
	}
	return list
}
