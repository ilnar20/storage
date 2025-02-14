package main

import (
	"fmt"
	"log"

	"github.com/ilnar20/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	f, err := st.Upload("test.txt", []byte("hello!"))
	if err != nil {
		log.Fatal(err)
	}

	retrivedFile, err := st.GetById(f.Id)
	if err != nil {
		log.Fatalf("error occured: %v", err)
	}
	fmt.Println(retrivedFile)
}
