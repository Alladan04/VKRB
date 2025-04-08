package utils

func Uint8ToInt(labirint [][]uint8) [][]int {
	res := make([][]int, len(labirint))
	for i, v := range labirint {
		for _, v2 := range v {
			res[i] = append(res[i], int(v2))
		}
	}

	return res
}
