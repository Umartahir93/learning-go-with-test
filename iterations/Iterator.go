package iterations

const iterationLoop = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < iterationLoop; i++ {
		repeated += character
	}
	return repeated

}
