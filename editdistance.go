package main

import "fmt"

const (
	minEditdistance = 0
)

type Row []int

// 实现一个求俩个字符串最小编辑距离
// str1 => str2 插入、删除、替换、交换的最小次数
func getEditDistance(str1 string, str2 string) int {
	rune1, rune2 := []rune(str1), []rune(str2)
	return minDistance(rune1, rune2)
}

// 求对应子串之间的编辑编辑距离
func minDistance(rune1 []rune, rune2 []rune) int {
	len1, len2 := len(rune1), len(rune2)
	// 初始化矩阵
	matrix := initMatrix(len1, len2)
	// 矩阵对应子串的编辑距离
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			minValue := min(matrix[i-1][j]+1, matrix[i][j-1]+1)
			if rune1[i-1] == rune2[j-1] {
				minValue = min(minValue, matrix[i-1][j-1])
			} else {
				minValue = min(minValue, matrix[i-1][j-1]+1)
			}

			// 对相邻字符交换位置的判断处理
			if i > 1 && j > 1 {
				if rune1[i-1] == rune2[j-2] && rune1[i-2] == rune2[j-1] {
					minValue = min(minValue, matrix[i-2][j-2]+1)
				}
			}
			matrix[i][j] = minValue
		}
	}
	fmt.Println(matrix)
	return matrix[len1][len2]
}

// 初始化一个矩阵
func initMatrix(len1, len2 int) []Row {
	matrix := make([]Row, len1+1)
	for i := 0; i <= len1; i++ {
		matrix[i] = make(Row, len2+1)
		for j := 0; j <= len2; j++ {
			matrix[i][j] = minEditdistance
			if i == 0 || j == 0 {
				matrix[i][j] = max(i, j)
			}
		}
	}
	return matrix
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	str1 := "sailn"
	str2 := "failing"
	// str1 := "哈哈哈哈"
	// str2 := "😆哈哈"
	dis := getEditDistance(str1, str2)
	fmt.Printf("字符串str1变换到字符串str2最少变换次数为:[%d]\n", dis)
}
