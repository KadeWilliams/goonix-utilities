package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func cat(path string) {
	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f);

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func ls(path string) {
	if path == "." {
		path, _ = os.Getwd()
	}
	if files, err := os.ReadDir(path); err == nil {
		fmt.Println(files)
		for _, file := range files {
			fmt.Printf("%s %v\n",file.Name(), file.IsDir())
		}
	}
}

func grep(value, path string) {
	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f);

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, value) {
			fmt.Println(line)
		}
	}
}

const pageSize = 24
func less(path string) {
	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f);
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	currentLine := 0
	for {
		clearScreen()
		for i := currentLine; i < currentLine+pageSize && i < len(lines); i++ {
			fmt.Println(lines[i])
		}
		fmt.Printf("\n-- More -- (Line %d of %d)", currentLine + 1, len(lines))

		input := waitForInput()

		switch input {
		case "q":
			return
		case " ":
			currentLine += pageSize
			if currentLine >= len(lines) {
				currentLine = len(lines) - 1
			}
		case "b": 
			currentLine -= pageSize 
			if currentLine < 0 {
				currentLine = 0
			}
		}
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func waitForInput() string {
	reader := bufio.NewReader(os.Stdin) 
	input, _ := reader.ReadString('\n')
	return string(input[0])
}

func sort(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sortedLines := []string{}

	for i := range lines {
		fmt.Printf("Outer: %s\n", lines[i])
		for j := range lines {
			fmt.Printf("\tInner:%s\n", lines[j])
		}
		

		// if len(sortedLines) < 1 {
		// 	sortedLines = append(sortedLines, lines[i])
		// 	continue
		// }

		// // TODO need to make it work when the characters are the same in pos 1 
		// for j := range sortedLines {
		// 	fmt.Printf("comparing %s to %s\n", lines[i], sortedLines[j])
		// 	if sortedLines[j] > lines[i] {
		// 		fmt.Printf("Inserting %s @ %d\n", lines[j], j)
		// 		sortedLines = slices.Insert(sortedLines, j, lines[j])
		// 		// sortedLines = append([]string{lines[i]}, sortedLines...)
		// 		// continue
		// 	} else {
		// 		fmt.Printf("Inserting %s @ %d\n", lines[i], i)
		// 		sortedLines = slices.Insert(sortedLines, i, lines[i])
		// 		break
		// 	}
		// 	// fmt.Printf("Appending %s\n", lines[i])
		// 	// sortedLines = append(sortedLines, lines[i])
		// 	// break
		// }
		

		// for j := range sortedLines {
		// 	if lines[i] < sortedLines[j] {
		// 		sortedLines = append([]string{lines[i]}, sortedLines...)
		// 	} else {
		// 		sortedLines = append(sortedLines, lines[i])
		// 	}
		// }
	}
	fmt.Println(sortedLines)
	// for _, line := range lines {
	// 	for _, sortedLine := range sortedLines {
	// 		if line < sortedLine 
	// 	}
	// }

}

func wc(path string) {
	file1, err := os.Open(path)
	file2, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file1.Close()

	scanner := bufio.NewScanner(file1)

	wordCount := 0
	lineCount := 0
	charCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineSlice := strings.Split(line, " ")
		for _, line := range lineSlice {
			for range line {
				charCount++ 
			}
		}
		wordCount += len(lineSlice)
		lineCount++
	}

	data := make([]byte, 32*1024)
	byteCount, err := file2.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	

	fmt.Printf("%d %d %d %d\n", lineCount, byteCount, wordCount, charCount)
}


func main() {
	args := os.Args[1:]

	switch args[0] {
	case "cat": 
		cat(args[1])
	case "ls": 
		ls(args[1])
	case "grep": 
		searchValue := args[1]
		searchPath := args[2]
		grep(searchValue, searchPath)
	case "less":
		less(args[1])
	case "sort":
		sort(args[1])
	case "tail": 
		break
	case "wc": 
		wc(args[1])
	}

}