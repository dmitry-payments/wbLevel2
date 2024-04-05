package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Флаги командной строки
	fieldsStr := flag.String("f", "", "выбрать поля (колонки)")
	delimiter := flag.String("d", "", "использовать другой разделитель")
	separated := flag.Bool("s", false, "только строки с разделителем")
	flag.Parse()

	fields := map[int]struct{}{}

	for _, field := range strings.Split(*fieldsStr, ",") {
		field := strings.TrimSpace(field)
		if field == "" {
			continue
		}
		i, err := strconv.Atoi(field)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if i < 0 {
			fmt.Println("поля нумеруются с 1")
			os.Exit(1)
		}
		fields[i] = struct{}{}
	}

	if *delimiter == "" {
		*delimiter = "\t"
	}

	// Сканер для чтения ввода
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}
		columns := strings.Split(line, *delimiter)
		for i, val := range columns {
			if _, ok := fields[i+1]; !ok {
				continue
			}
			fmt.Printf("%s ", val)
		}
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении ввода:", err)
		os.Exit(1)
	}
}
