package main

import "fmt"

func main() {
array := make([][]int, 5)

// 创建二维Slice
for i := range array {
    subArray := make([]int, 4)
    for j := range subArray {
        subArray[j] = j + 1
    }
    array[i] = subArray
}

// 输出
for i := range array[1:4][2:] {
    for j := range array[i] {
        fmt.Printf("%v (%d, %d) ", array[i][j], i, j)
    }
    fmt.Println()
}
}
