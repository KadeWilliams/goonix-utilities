package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
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

func ls (path string) {
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

func wc(path string) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	fileInfo.Size()

	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}
	r, _ := os.Open(path)
	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			fmt.Println(count)
			break
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
		break 
	case "sort":
		break
	case "tail": 
		break
	case "wc": 
		wc(args[1])
	}

}