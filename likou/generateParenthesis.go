package likou

//回溯要点
//1.选择
//2.约束条件
//3.目标
func GenerateParenthesis(n int) []string {
	var result []string
	var dfs func(l, r int, str string)
	dfs = func(l, r int, str string) {
		//递归出口
		if len(str) == 2*n {
			result = append(result, str)
			return
		}
		//约束条件
		if l > 0 {
			//目标
			dfs(l-1, r, str + "(")
		}
		//约束条件
		if l < r {
			//目标
			//不能直接改变str
			dfs(l, r-1, str + ")")
		}

	}
	dfs(n, n, "")
	return result
}
