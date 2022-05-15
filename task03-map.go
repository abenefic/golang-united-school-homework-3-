package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	var tmp []int
	for key := range input {
		tmp = append(tmp, key)
	}
	sort.Ints(tmp)
	for _, key := range tmp {
		result = append(result, input[key])
	}
	return
}
