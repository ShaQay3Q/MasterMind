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

func compare(answer combination, guess combination) hint {

	res := compareSameIndexes(answer, guess)

	return res
}

func compareSameIndexes(answer combination, guess combination) hint {
	var res hint

	for i := range answer {
		if answer[i] == guess[i] {
			res.black += 1
		}
	}

	return res
}
