package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var inputReader *bufio.Reader

func read() (string, string) {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Print("input length:")
	inputLength, err := inputReader.ReadString('\n')
	if err != nil {
		return "", ""
	}

	fmt.Print("input slice and separation with space:")
	inputSlice, err := inputReader.ReadString('\n')
	if err != nil {
		return "", ""
	}
	return inputLength, inputSlice
}

func parseInput(l, s string) (int, []int, error) {
	var sli []int

	str := l[:len(l)-1]
	length, err := string2Int(str)
	if err != nil {
		return 0, sli, err
	}

	str = s[:len(s)-1]

	inputSlice := strings.Split(str, " ")
	for _, is := range inputSlice {
		n, err := string2Int(is)
		if err != nil {
			return 0, sli, err
		}
		sli = append(sli, n)
	}

	return length, sli, nil
}

func string2Int(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Errorf("Parse string error: %v", err)
		return 0, err
	}

	return n, nil
}

func compare(n int, ns []int) bool {
	sort.Ints(ns)

	if n != len(ns) {
		return false
	}

	for i := 0; i < n-2; i++ {
		if ns[i+1]-ns[i] != ns[i+2]-ns[i+1] {
			return false
		}
	}

	return true
}

func main() {
	length, nslice, err := parseInput(read())
	if err != nil {
		panic(err)
	}
	if ok := compare(length, nslice); !ok {
		fmt.Print("Impossible \n")
	} else {
		fmt.Print("Possible \n")
	}
}
