package main

import (
	"log"
	"os"
	"testing"
)

// Создание тестового файла
func makeFile(fileName string){
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err = f.Write([]byte("row1\nrow2")); err != nil{
		log.Fatal(err)
	}
}

// Удаление тестового файла
func delFile(fileName string){
	if err := os.Remove(fileName); err != nil {
		log.Fatal(err)
	}
}


func TestGetUrls(t *testing.T) {
	var(
		testFile = "urls.txt"
		result = []string{"row1", "row2"}
	)

	makeFile(testFile)
	defer delFile(testFile)

	g, err := GetUrls(testFile)
	if err != nil {
		t.Fatal(err)
	}
	if g[0] != result[0] || g[1] != result[1] || len(g) != len(result){
		t.Fatalf("Ошибка. Ожидается %v, получено: %v\n", result, g)
	}
}
