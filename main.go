package main

import (
	"./models"
	"fmt"
)

func main() {
	str1 := "sailn"
	str2 := "failing"
	// str1 := "哈哈哈哈"
	// str2 := "😆哈哈"
	dis := models.GetEditDistance(str1, str2)
	fmt.Printf("字符串str1变换到字符串str2最少变换次数为:[%d]\n", dis)
}
