package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	urlPtr := flag.String("url", "", "URL сайта для загрузки")
	flag.Parse()

	if *urlPtr == "" {
		fmt.Println("Необходимо указать URL сайта для загрузки")
		return
	}

	if !strings.HasPrefix(*urlPtr, "http://") && !strings.HasPrefix(*urlPtr, "https://") {
		*urlPtr = "http://" + *urlPtr
	}

	response, err := http.Get(*urlPtr)
	if err != nil {
		fmt.Println("Ошибка при получении данных сайта:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении данных сайта:", err)
		return
	}

	fileName := "index.html"
	if strings.HasSuffix(*urlPtr, "/") {
		fileName = strings.TrimRight(strings.TrimPrefix(*urlPtr, "http://"), "/") + ".html"
	}
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(body)
	if err != nil {
		fmt.Println("Ошибка при записи данных в файл:", err)
		return
	}

	fmt.Printf("Сайт успешно загружен и сохранен в файл %s\n", fileName)
}
