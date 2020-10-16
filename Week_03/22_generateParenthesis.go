package Week_01

func generateParenthesis(n int) []string {
	res := new([]string)

	generate(res, "", 0, 0, n)

	return *res;
}

func generate(res *[]string, s string, open int, close int, max int)  {
	if open == max && close == max {
		*res = append(*res, s)
		return
	}

	if open < max {
		generate(res, s + "(", open + 1, close, max)
	}

	if close < open {
		generate(res, s + ")", open, close + 1, max)
	}
}