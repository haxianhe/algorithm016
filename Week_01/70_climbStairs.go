package Week_01

func climbStairs(n int) int {
	tmp1,tmp2,result := 0,0,1

	for i := 1; i<= n; i++ {
		tmp1 = tmp2
		tmp2 = result
		result = tmp1 + tmp2
	}

	return result
}

//暴力？基本情况
//1:1
//2:2:
//3:f(1)+f(2)
//4:f(2)+f(3)
//...
//n:f(n-2)+f(n-1)
//找 最近 重复子问题
