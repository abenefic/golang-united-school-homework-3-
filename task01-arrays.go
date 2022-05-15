package homework

func average(input [15]float32) (result float32) {
	//Place your code here
	for _, i := range input {
		result += i
	}
	result = result / 15
	return
}
