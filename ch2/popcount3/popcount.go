package popcount3 // import "github.com/AntiBargu/gopl/ch2/popcount3"

// PopCount 返回x的种群统计（置位的个数）
func PopCount(x uint64) int {
	var rslt int

	for x != 0 {
		rslt++
		x = x & (x - 1)
	}

	return rslt
}
