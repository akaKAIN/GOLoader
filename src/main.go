package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main(){
	url := os.Args[1]
	format := "mp4"
	if err := LoadAndSaveFile(url, format); err != nil{
		log.Printf("Ошибка загрузчика по ссылке %s. Error: %s", url, err)
	}
}


func LoadAndSaveFile(url string, format string) error{
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	newFile, err := os.Create("fileName." + format)
	if err != nil{
		return err
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, resp.Body)
	return err
}