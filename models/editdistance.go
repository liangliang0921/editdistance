package models

import (
	"../utils"
	"fmt"
)

const (
	minEditdistance = 0
)

type Row []int

type GetEditDistanceResqData struct {
	NewString    string `json:"newstring"`
	OldString    string `json:"oldstring"`
	EditDistance int    `json:"editdistance"`
}

// 实现一个求俩个字符串最小编辑距离
// str1 => str2 插入、删除、替换、交换的最小次数
func GetEditDistance(str1 string, str2 string) *GetEditDistanceResqData {
	rune1, rune2 := []rune(str1), []rune(str2)
	editDistanceData := GetEditDistanceResqData{
		NewString:    str1,
		OldString:    str2,
		EditDistance: minDistance(rune1, rune2),
	}
	return &editDistanceData
}

// 求对应子串之间的编辑编辑距离
func minDistance(rune1 []rune, rune2 []rune) int {
	len1, len2 := len(rune1), len(rune2)
	// 初始化矩阵
	matrix := initMatrix(len1, len2)
	// 矩阵对应子串的编辑距离
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			minValue := utils.Min(matrix[i-1][j]+1, matrix[i][j-1]+1)
			if rune1[i-1] == rune2[j-1] {
				minValue = utils.Min(minValue, matrix[i-1][j-1])
			} else {
				minValue = utils.Min(minValue, matrix[i-1][j-1]+1)
			}

			// 对相邻字符交换位置的判断处理
			if i > 1 && j > 1 {
				if rune1[i-1] == rune2[j-2] && rune1[i-2] == rune2[j-1] {
					minValue = utils.Min(minValue, matrix[i-2][j-2]+1)
				}
			}
			matrix[i][j] = minValue
		}
	}
	// fmt.Println(matrix)
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
				matrix[i][j] = utils.Max(i, j)
			}
		}
	}
	return matrix
}
