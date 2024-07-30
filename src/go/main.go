package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("💀 Please provide valid arguments\n>> program <destination_file> <currency-sign> <separator>")
		os.Exit(1)
	}

	destFilePath := os.Args[1]
	currencySign := os.Args[2]
	separator := os.Args[3]

	destFile, err := os.Create(destFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating destination file: %v\n", err)
		os.Exit(1)
	}
	defer destFile.Close()

	// Writer buffer
	destFileWriter := bufio.NewWriter(destFile)
	defer destFileWriter.Flush()

	// Reader buffer
	stdinReader := bufio.NewReader(os.Stdin)

	for {
		num, err := stdinReader.ReadString(' ')
		isProgramEnd := false
		if err != nil {
			if err == io.EOF {
				isProgramEnd = true
			} else {
				fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
				os.Exit(1)
			}
		}
		processedByte := ""

		trimNum := strings.TrimSpace(num)
		if _, isNotValidNum := strconv.Atoi(trimNum); isNotValidNum != nil {
			if isProgramEnd {
				break
			}
			continue
		}

		separatedNumber := separateNumber(trimNum, separator)
		processedByte += fmt.Sprintf("%s%s ", currencySign, separatedNumber)
		_, err = destFileWriter.WriteString(processedByte)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to %s: %v\n", destFilePath, err)
			os.Exit(1)
		}

		if isProgramEnd {
			break
		}
	}

	fmt.Println("Processing complete. Output written to:", destFilePath)
}

func separateNumber(number, separator string) string {
	var result []string
	for i := len(number); i > 0; i -= 3 {
		start := i - 3
		if start < 0 {
			start = 0
		}
		result = append([]string{number[start:i]}, result...)
	}

	return strings.Join(result, separator)
}
