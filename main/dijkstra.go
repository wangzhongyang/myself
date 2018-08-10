package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 初始化点的个数
	str := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	// 生成矩阵及各边的权重（随机数，非负正整数）
	vertexArray, edgesMap := Matrix(str)
	fmt.Println(vertexArray, edgesMap)
	later := make(map[string]int)     // 记录已经过点
	minLength, key, length := 0, 0, 0 // 路长
	minPath, start, end := str[0], "", ""
	for {
		if key == 0 {
			start = str[0]
		} else {
			start = end
		}

		end, length = MinVertex(start, vertexArray, edgesMap, later)
		fmt.Println("----------", start, end, length)
		fmt.Println(later)
		if _, ok := later[end]; ok {
			break
		} else {
			later[end] = 1
			minLength += length
			minPath += end
		}
		key++
	}
	fmt.Println(minPath, minLength)
}

// 生成矩阵
func Matrix(args []string) (map[string][]string, map[string]int) {
	vertexArray := map[string][]string{} // 顶点map
	edgesMap := make(map[string]int)     // 边map
	for key1, val1 := range args {
		vertexArr := []string{}
		for key2, val2 := range args {
			if key2 != 0 && key2 != key1 {
				vertexArr = append(vertexArr, val2)
				edgesMap[val1+val2] = rand.Intn(20)
			}
		}
		vertexArray[val1] = vertexArr
	}
	return vertexArray, edgesMap
}

// 获取距离start点最近的点
func MinVertex(start string, vertexArray map[string][]string, edgesMap, later map[string]int) (string, int) {
	nextArr := vertexArray[start]
	minLen := 9999999
	end := nextArr[0]
	for _, v := range nextArr {
		if _, ok := later[v]; !ok && minLen > edgesMap[start+v] {
			minLen = edgesMap[start+v]
			end = v
		}
	}
	return end, minLen
}
