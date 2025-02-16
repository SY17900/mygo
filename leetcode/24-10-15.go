package leetcode

func findheight(topblue bool, red int, blue int) int {
	height := 0
	if topblue {
		if blue == 0 {
			return 0
		}
		for {
			if (height+1)%2 == 1 {
				blue -= height + 1
			} else {
				red -= height + 1
			}
			if blue < 0 || red < 0 {
				return height
			}
			height++
		}
	} else {
		if red == 0 {
			return 0
		}
		for {
			if (height+1)%2 == 1 {
				red -= height + 1
			} else {
				blue -= height + 1
			}
			if blue < 0 || red < 0 {
				return height
			}
			height++
		}
	}
}

func MaxHeightOfTriangle(red int, blue int) int {
	return max(findheight(true, red, blue), findheight(false, red, blue))
}
