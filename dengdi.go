package main

import "fmt"

func main(){
	fmt.Println("输入N：")
       fmt.Println(outOrder(Arr()))
}

func outOrder(trainsNums []string) []string {
	COUNT := len(trainsNums)
	//检查
	if COUNT == 0 || COUNT > 12 {
		panic("Illegal argument. trainsNums size must between 1 and 9.")
	}

	//如果只有一个数，则直接返回
	if COUNT == 1 {
		return []string{trainsNums[0]}
	}

	//否则，将最后一个数插入到前面的排列数中的所有位置（递归）

	return insert(outOrder(trainsNums[:COUNT - 1]), trainsNums[COUNT - 1])
}
func insert(res []string, inserts string )[]string{
	result := make([]string, len(res)*(len(res[0])+1))
	index := 0
	for _, v := range res {
		for i := 0; i < len(v); i++ {
			//在v的每一个元素前面插入
			result[index] = v[:i] + inserts + v[i:]
			// fmt.Print("head:", result[index])
			index++
		}

		//在v最后面插入
		result[index] = v + inserts
		// fmt.Print("tail:", result[index])
		index++
	}
	// fmt.Print(result, ">>")
	return result
}
func Arr()[]string {
	var Input int
	fmt.Scanln(&Input)
	arr := [19]string{"1","2","3","4","5","6","7","8","9","10","11","12"}
	slice := arr[:Input]
	return slice

}