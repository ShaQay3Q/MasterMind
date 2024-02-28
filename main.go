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

func compareIndexes(answer combination, guess combination) hint {

	var res hint
	list := []entry{}

	for i := range answer {
		a := sliceIt(answer, i)
		for j := range guess {
			g := sliceIt(guess, j)
			switch {
			case compare(a, g):
				switch {
				case i == j:
					res.black += 1
					break
				case i != j:
					list = addToList(a, j, list)
					res.white = len(addToList(a, j, list))
					break
				}
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

func addToList(col color, digit int, list []entry) []entry {

	newEntry := entry{col, digit}
	check := checkForColorAndIndex(newEntry, list)

	switch {
	case len(list) == 0:
		list = append(list, newEntry)
		return list
	case len(list) != 0:
		switch {
		case check == true:
			break
		case check != true:
			list = append(list, newEntry)
			break
		}
	}
	return list
}

func checkForColorAndIndex(newEntry entry, list []entry) bool {
	for i := 0; i < len(list); i++ {
		if newEntry.color == list[i].color && newEntry.digit == list[i].digit {
			return true
		}
	}
	return false
}
