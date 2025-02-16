package leetcode

func GenerateKey(x int, y int, z int) (ans int) {
	for times := 0; x > 0 && y > 0 && z > 0; times *= 10 {
		ans += min(x%10, y%10, z%10) * times
		x /= 10
		y /= 10
		z /= 10
	}
	return
}
