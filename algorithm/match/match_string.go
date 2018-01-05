package match

import (
	"bufio"
	"fmt"
	"os"
)

var (
	inputReader *bufio.Reader
	buffer      []byte
	longString  string
)

const (
	commonMultiCharacter = "*"
	commonCharacter      = "?"
	defaultSize          = 65336
)

func init() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	buffer = make([]byte, defaultSize)

	num, err := file.Read(buffer)
	if err != nil {
		panic(err)
	}

	if num > 0 {
		longString = string(buffer[:num])
	}
}

func readFromInput() string {
	fmt.Print("input match string:")

	inputReader = bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		return ""
	}

	return input[:len(input)-1]
}

func handleInput(input string) {
	n := len(input)
	if n < 1 {
		return
	}

	if n == 1 {
		handleSingleCharacter(input, 0)
	}

}

func handleSingleCharacter(char string, start int) string {
	if char == commonMultiCharacter {
		fmt.Printf("match string is:%s \n", longString[start:])
		return longString[start:]
	}

	if char == commonCharacter {
		fmt.Printf("match string is:%s \n", longString[start:start+1])
		return longString[start : start+1]
	}

	asc := uint8([]byte(char)[0])

	if asc >= 65 && asc <= 90 || asc >= 97 && asc <= 122 {
		fmt.Printf("you just input one character:%s \n", char)
	}

}
