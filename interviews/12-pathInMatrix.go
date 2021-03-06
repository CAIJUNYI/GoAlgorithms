package interviews

// 请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。
//
// [["a","b","c","e"],
// ["s","f","c","s"],
// ["a","d","e","e"]]
//
// 但矩阵中不包含字符串“abfb”的路径，因为字符串的第一个字符b占据了矩阵中的第一行第二个格子之后，路径不能再次进入这个格子。
//
//
// 示例 1：
// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
// 输出：true
//
// 示例 2：
// 输入：board = [["a","b"],["c","d"]], word = "abcd"
// 输出：false
// 提示：
//
// 1 <= board.length <= 200
// 1 <= board[i].length <= 200

func exist(board [][]byte, word string) bool {
	if len(board) < 1 || len(board[0]) < 1 {
		return false
	}

	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var res bool
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res = res || existPath(board, []byte(word), visited, i, j)
		}
	}

	return res
}

func existPath(board [][]byte, chars []byte, visited [][]bool, i, j int) bool {
	// the end
	if len(chars) == 0 {
		return true
	}

	// check board
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}

	// visited
	if visited[i][j] {
		return false
	}

	c := chars[0]
	if board[i][j] != c {
		return false
	}

	visited[i][j] = true
	res := existPath(board, chars[1:], visited, i-1, j) ||
		existPath(board, chars[1:], visited, i+1, j) ||
		existPath(board, chars[1:], visited, i, j-1) ||
		existPath(board, chars[1:], visited, i, j+1)

	visited[i][j] = false

	return res
}
