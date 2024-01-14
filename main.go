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

func compare(answer combination, guess combination) hint {
	var res hint
	res.black = 4
	res.white = 0

	return res
}
