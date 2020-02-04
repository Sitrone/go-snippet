package leetcode

type sudoSet map[byte]struct{}

func (s *sudoSet) Add(val byte) {
	(*s)[val] = struct{}{}
}

func (s *sudoSet) Contains(val byte) bool {
	_, ok := (*s)[val]
	return ok
}

func (s *sudoSet) Reset() {
	for k := range *s {
		delete(*s, k)
	}
}

func isValidSudoku(board [][]byte) bool {
	if len(board) == 0 || len(board) != 9 || len(board[0]) != 9 {
		return false
	}
	const spaceNum = byte('.')

	row, col := len(board), len(board[0])

	var colMap, rowMap = &sudoSet{}, &sudoSet{}
	for i := 0; i < row; i++ {
		rowMap.Reset() // 直接分配map，性能高，内存占用高，重复使用map，性能下降，内存占用降低
		colMap.Reset()

		for j := 0; j < col; j++ {
			if board[i][j] != spaceNum {
				if rowMap.Contains(board[i][j]) {
					return false
				} else {
					rowMap.Add(board[i][j])
				}
			}

			if board[j][i] != spaceNum {
				if colMap.Contains(board[j][i]) {
					return false
				} else {
					colMap.Add(board[j][i])
				}
			}
		}
	}

	var m = &sudoSet{}
	for i := 0; i < row; i += 3 {
		for j := 0; j < col; j += 3 {
			m.Reset()
			for ik := i; ik < i+3; ik++ {
				for jk := j; jk < j+3; jk++ {
					if board[ik][jk] != spaceNum {
						if m.Contains(board[ik][jk]) {
							return false
						} else {
							m.Add(board[ik][jk])
						}
					}
				}
			}
		}
	}

	return true
}
