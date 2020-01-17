package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

const format = "mp4"

func main() {
	var (
		wg       sync.WaitGroup
		urls     []string
		fileName string
		counter int
	)
	switch len(os.Args) {
	case 1:
		fileName = "urls.txt"
	case 2:
		fileName = os.Args[1]
	default:
		log.Println("Неверно указаны параметры команды.")
		os.Exit(1)
	}

	urls, err := GetUrls(fileName)
	fmt.Printf("Получено ссылок из файла: %d\n", len(urls))
	if err != nil {
		log.Fatalf("Ошибка чтения файла. Файл %q не найден.\n", fileName)
	}
	for _, url := range urls {
		if url != ""{
			counter++
			wg.Add(1)
			go LoadAndSaveFile(url, counter, format, &wg)
		}

	}
	wg.Wait()
	log.Println("Все ссылки прошли загрузку.")
}

func GetUrls(fileName string) ([]string, error) {
	var urls []string

	fileBody, err := os.Open(fileName)
	if err != nil {
		err = fmt.Errorf("Ошибка открытия файла (%s)\n", err)
		return nil, err
	}
	defer fileBody.Close()

	r := bufio.NewReader(fileBody)
	for true {
		byteRow, _, _ := r.ReadLine()
		if len(byteRow) == 0 {
			break
		}
		urls = append(urls, string(byteRow))

	}
	return urls, nil
}

func LoadAndSaveFile(url string, index int, format string, wg *sync.WaitGroup) {
	var strIndex, fileName string
	defer wg.Done()

	strIndex = strconv.Itoa(index)
	fileName = "video_" + strIndex + "." + format
	if url == ""{
		err := fmt.Errorf("Получен пустой адресс ссылки.\n")
		log.Println(err)
	}
	resp, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("Ошибка получения ответа от сервера (%v)\n", err)
		log.Println(err)
	}
	defer resp.Body.Close()
	log.Printf("Скачивание файла %q по ссылке: %s\n", fileName, url)

	newFile, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, resp.Body)
	log.Printf("Файл %q успешно сохранен.\n", fileName)
}
