package Week_01

func moveZeroes(nums []int)  {
	if len(nums) == 0 || nums == nil {
		return
	}

	var j = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] !=0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}
//1.遍历，将所有非零元素放入新数组，记录零的个数，遍历后补零
//2.遍历，如果非零将当前指针元素放入慢指针位置，然后向后移动当前指针和慢指针，遍历将慢指针及之后的元素置为0