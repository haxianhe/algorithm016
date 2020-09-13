package Week_01

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	water := 0
	for ;i<j; {
		water = max(water,(j - i) * min(height[i], height[j]))
		if height[i] < height[j] {
			i++
		}else {
			j--
		}
	}
	return water
}

func max(x,y int) int {
	if x <= y {
		return  y
	}else {
		return  x
	}
}

func min(x,y int) int  {
	if x >= y {
		return  y
	}else {
		return  x
	}
}

//1.water = max(water,(j-i)*min(height[i],height[j])),并不断向内移动左右指针对应值小的那个
