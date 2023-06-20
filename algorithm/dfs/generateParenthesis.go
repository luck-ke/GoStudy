package dfs

// Parenthesis 生成有效括号对
func Parenthesis(n int) []string {
	result := make([]string, 0)
	dfs(&result, "", n*2, n*2)
	return result
}

func dfs(result *[]string, str string, n, target int) {
	if n == 0 {
		*result = append(*result, str)
		return
	}
	dfs(result, str+"(", n-1, target)
	if n != target {
		dfs(result, str+")", n-1, target)
	}

}
