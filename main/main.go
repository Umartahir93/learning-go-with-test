package main

import "fmt"

const spanish = "Spanish"
const french = "French"

const helloEnglishPrefix = "Hello, "
const holaEnglishPrefix = "Hola, "
const bonjourFrenchPrefix = "Bonjour, "

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = bonjourFrenchPrefix
	case spanish:
		prefix = holaEnglishPrefix
	default:
		prefix = helloEnglishPrefix
	}

	return
}

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums

}

func main() {
	fmt.Println(Hello("Umar", ""))
}
