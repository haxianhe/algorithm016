package Week_02

/*
1.暴力
1.1 先遍历第一个字符串统计每个字母出现的次数
1.2 后遍历第二个字符串统计每个字母出现的次数
1.3 最后遍历统计数组 比较两个数组每个字母出现的次数是否一致
 */

/*
2.暴力
2.1 先遍历第一个字符串统计每个字母出现的次数
3.2 后遍历第二个字符串每次在统计数组中减一
3.3 判断统计数组是否为空，为空则返回true 否则返回false
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	strMap := make(map[string]int)

	for i := 0; i < len(s); i++ {
		strMap[string(s[i])]++
	}

	for i := 0; i < len(t); i++ {
		strMap[string(t[i])]--

		if strMap[string(t[i])] < 0 {
			return false;
		}
	}

	for _, v := range strMap{
		if v!= 0 {
			return false;
		}
	}

	return true
}