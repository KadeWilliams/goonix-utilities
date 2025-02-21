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
		fmt.Println("IsDirectory\tObject Name")
		for _, file := range files {
			fmt.Printf("%v\t\t%s\n", file.IsDir(), file.Name())
		}
	}
}

func grep(value, path string) {
	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f);

	line := 1
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, value) {
			fmt.Printf("[line %d] %s\n", line, text)
		}
		line++
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

	for i := range lines {
		for j := range lines {
			if lines[i] < lines[j] {
				flipVal := lines[i]
				lines[i] = lines[j]
				lines[j] = flipVal
			}
		}
	}
	fmt.Println(lines)
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

func tail(path string) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	const maxLines = 10
	ringBuffer := make([]string, maxLines)

	index := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ringBuffer[index] = scanner.Text()
		index = (index + 1) % maxLines
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for i := 0; i < maxLines; i++ {
		line := ringBuffer[(index+i)%maxLines]
		if line != "" {
			fmt.Println(line)
		}
	}
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
		tail(args[1])
	case "wc": 
		wc(args[1])
	}

}