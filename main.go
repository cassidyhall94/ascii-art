package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		return
	}

	// Writing arguments in a single string
	string := os.Args[1]
	for _, v := range os.Args[2:] {
		string += " " + v
	}

	// Checking whether str contain "\n" or not ---> executing the ascii-art
	previous := 'a'
	manylines := false
	for _, v := range string {
		if v == 'n' && previous == '\\' {
			manylines = true
		}
		previous = v
	}
	// Writing text line by line into result
	result := ""
	if manylines {
		args := strings.Split(string, "\\n")
		for _, word := range args {
			for i := 0; i < 8; i++ {
				for _, char := range word {
					result += ReturnLine(1 + int(char-' ')*9 + i)
				}
				fmt.Println(result)
				result = ""
			}
		}

	} else {
		for i := 0; i < 8; i++ {
			for _, char := range string {
				result += ReturnLine(1 + int(char-' ')*9 + i)
			}
			fmt.Println(result)
			result = ""
		}
	}
}

func ReturnLine(num int) string {
	string := ""
	f, e := os.Open("standard.txt")
	if e != nil {
		fmt.Println(e.Error())
		os.Exit(0)
	}
	defer f.Close()

	f.Seek(0, 0)
	content := bufio.NewReader(f)
	for i := 0; i < num; i++ {
		string, _ = content.ReadString('\n')
	}
	string = strings.TrimSuffix(string, "\n")
	return string
}
