package arraylib

func GetLargestElementAndIndex(arr []int) (int, int) {
	var max_el int = 0
	var max_idx int = -1
	for i, v := range arr {
		if v > max_el {
			max_el = v
			max_idx = i
		}
	}
	return max_idx, max_el
}

func Sum(arr []int) int{
	sum := 0
	for _, v := range arr{
		sum += v
	}
	return sum
}