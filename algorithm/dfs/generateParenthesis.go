package dfs

// Parenthesis 生成有效括号对
func Parenthesis(n int) []string {
	result := make([]string, 0)
	dfs(&result, "", n, n)
	return result
}

func dfs(result *[]string, str string, l, r int) {
	if r < l {
		return
	}
	if l == 0 && r == 0 {
		*result = append(*result, str)
		return
	}
	if l > 0 {
		dfs(result, str+"(", l-1, r)
	}
	if str != "" && r > 0 {
		dfs(result, str+")", l, r-1)
	}
}
