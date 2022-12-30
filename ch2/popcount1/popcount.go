package popcount1 // import "github.com/AntiBargu/gopl/ch2/popcount1"

// pc[i] 是i的种群统计
var pc [256]byte

// 计算0-255（1个字节）所有值的置位数
func init() {
	// for i, _ := range pc {
	// 简要写法
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount 返回x的种群统计（置位的个数）
func PopCount(x uint64) int {
	var rslt int

	for i := 0; i < 8; i++ {
		rslt += int(pc[byte(x>>(i*8))])
	}

	return rslt
}
