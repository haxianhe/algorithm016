package Week_01

/**
解法1：暴力递归 O(2^n)
 */
//func climbStairs(n int) int {
//	if n <= 2 {
//		return n
//	}
//	return climbStairs(n - 1) + climbStairs(n -2)
//}

/**
解法2：动态规划法 O(n)
f(1) = 1
f(2) = 2
·
·
f(n) = f(n-1) + f(n-2)
 */
//func climbStairs(n int) int {
//	tmp1,tmp2,result := 0,0,1
//
//	for i := 1; i<= n; i++ {
//		tmp1 = tmp2
//		tmp2 = result
//		result = tmp1 + tmp2
//	}
//
//	return result
//}


/**
解法3：递归+缓存 O(n)
这是一个斐波那契数列问题，首先想到的是可以用递归来解决，但是简单的用递归解决会存在性能问题，那么除此之外还可以做什么呢？那就是加缓存
 */
//func climbStairs(n int) int {
//	catch := make(map[int]int)
//	return helper(n, catch)
//}
//
//func helper(n int, catch map[int]int) int {
//	//递归结束条件
//	if n <= 2 {
//		return n
//	}
//
//	//处理当前层逻辑
//	//进入下一层
//	ret := 0
//	if catch[n - 1] != 0 {
//		ret += catch[n - 1]
//	}else {
//		ret += helper(n - 1, catch)
//		catch[n - 1] = ret
//	}
//
//	if catch[n - 2] != 0 {
//		ret += catch[n - 2]
//	}else {
//		ret += helper(n - 2, catch)
//		catch[n - 2] = ret
//	}
//
//	//清理当前层
//	return ret
//}
