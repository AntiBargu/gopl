package popcount2 // import "github.com/rucetc/gopl/ch2/popcount2"

// PopCount 返回x的种群统计（置位的个数）
func PopCount(x uint64) int {
	var rslt int

	for i := 0; i < 64; i++ {
		if (x & (1 << i)) != 0 {
			rslt++
		}
	}

	return rslt
}
