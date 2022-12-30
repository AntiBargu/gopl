package popcount // import "github.com/AntiBargu/gopl/ch2/popcount"

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
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
