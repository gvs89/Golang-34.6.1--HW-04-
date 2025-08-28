package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: program <input_file> <output_file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Открываем входной файл
	inFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer inFile.Close()

	// Создаём или очищаем выходной файл
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer outFile.Close()

	// Буферизованный писатель
	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	// Регулярное выражение для поиска выражений
	re := regexp.MustCompile(`(\d+)([+-])(\d+)=\?`)
	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)

		if matches != nil {
			// Парсим числа
			a, _ := strconv.Atoi(matches[1])
			b, _ := strconv.Atoi(matches[3])
			op := matches[2]

			// Вычисляем результат
			var result int
			switch op {
			case "+":
				result = a + b
			case "-":
				result = a - b
			}

			// Формируем строку результата
			resultStr := fmt.Sprintf("%s%s%s=%d\n", matches[1], op, matches[3], result)
			writer.WriteString(resultStr)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
	}
}
