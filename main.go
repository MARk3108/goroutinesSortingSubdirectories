package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type FileInfo struct {
	Path string
	Size int64
}

func main() {
	dir := ""
	if len(os.Args) > 1 {
		dir = os.Args[1]
	} else {
		log.Fatal("Usage: go run . [directory_path]")
	}

	filesChan := make(chan FileInfo)
	var wg sync.WaitGroup

	go func() {
		wg.Wait()
		close(filesChan)
	}()

	wg.Add(1)
	go walkDir(dir, filesChan, &wg)

	var files []FileInfo
	for file := range filesChan {
		files = append(files, file)
	}

	sortFilesBySize(files)

	for _, file := range files {
		fmt.Printf("%s - %d bytes\n", file.Path, file.Size)
	}
}
