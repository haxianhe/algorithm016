package Week_01

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	m := make(map[int]int)

	//for i := 0; i < len(nums); i++ {
	//	j, ok := m[target - nums[i]]
	//	if ok != false {
	//		return []int{j, i}
	//	}
	//	m[nums[i]] = i
	//}

	for i, num := range nums {
		if j,ok := m[target - num]; ok != false {
			return []int{j, i}
		}
		m[nums[i]] = i
	}

	return nil
}

//1.暴力解法：列出所有解法 如果和为目标值返回相应下标
//2.使用hash表，遍历数组，每次判断target-当前值是否在hash表中，如果在返回相应值，如果不在将key=>value存入hash表中