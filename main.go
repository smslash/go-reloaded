package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"git/ssengerb/go-reloaded/goreloaded"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error: enter exactly two *.txt files")
		return
	}

	if os.Args[2] == os.Args[1] {
		fmt.Println("Error: created the same file as the original")
		return
	}

	if len(os.Args[2]) < 4 || len(os.Args[1]) < 4 {
		fmt.Println("Error: second *.txt file format is not correct")
		return
	} else {
		if os.Args[1][len(os.Args[1])-4] != 46 || os.Args[1][len(os.Args[1])-3] != 116 || os.Args[1][len(os.Args[1])-2] != 120 || os.Args[1][len(os.Args[1])-1] != 116 {
			fmt.Println("Error: first file should be .txt format")
			return
		}
		if os.Args[2][len(os.Args[2])-4] != 46 || os.Args[2][len(os.Args[2])-3] != 116 || os.Args[2][len(os.Args[2])-2] != 120 || os.Args[2][len(os.Args[2])-1] != 116 {
			fmt.Println("Error: second file should be .txt format")
			return
		}
	}

	file, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	sample := string(data)
	if len(sample) == 0 {
		_, err = file.WriteString(sample)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		return
	}

	result := goreloaded.Article(sample)
	for i := 0; i < len(result); i++ {
		result = goreloaded.Modification(result)
		result = goreloaded.Punctuation(result)
		if result == " " {
			return
		}
	}
	result = goreloaded.Modification(result)
	result = goreloaded.Apostrophe(result)
	result = goreloaded.Modification(result)

	_, err = file.WriteString(result)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
