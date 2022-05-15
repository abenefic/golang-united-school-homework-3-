package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	result = make([]int64, len(input))

	for i := 0; i < len(input); i++ {
		result = append(result[:len(input)-i], input[i])
	}
	return
}
