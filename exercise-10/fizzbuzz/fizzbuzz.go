package fizzbuzz

import "strconv"

const (
	fizz = "Fizz"
	buzz = "Buzz"
)

type Translator func(int) string

var Fizz Translator = func(i int) string {
	if i%3 == 0 {
		return fizz
	}

	return ""
}

var Buzz Translator = func(i int) string {
	if i%5 == 0 {
		return buzz
	}

	return ""
}

func Number(i int) string {
	if i%3 != 0 && i%5 != 0 {
		return strconv.Itoa(i)
	}

	return ""
}

func FizzBuzz(f Translator, i int) string {
	return f(i)
}
