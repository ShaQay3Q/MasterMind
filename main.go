package main

func main() {

}

type color string

type combination [4]color

type hint struct {
	b, w int
}

func compare(answer combination, guess combination) hint {
	var res hint
	res.b = 4
	res.w = 0

	return res
}
