package _go

func TowSum(nums []int, target int) []int {
	res := make([]int, 2)
	tMap := make(map[int]int)

	for i, n := range nums {
		tMap[target - n] = i
	}

	for i, n := range nums {
		if v, ok := tMap[n]; ok && v != i {
			if i < v {
				res[0] = i
				res[1] = v
			} else {
				res[0] = v
				res[1] = i
			}
		}
	}

	return res
}
