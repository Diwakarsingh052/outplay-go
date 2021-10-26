package sum

func SumInt(vs []int) int {

	sum := 0
	for _, v := range vs {
		sum = sum + v

	}
	if vs == nil {
		return 1
	}

	return sum

}
