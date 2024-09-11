package main

import (
	"os"
	"sync"
)

func processFile(path string, fileInfo os.FileInfo, filesChan chan FileInfo, wg *sync.WaitGroup) {
	defer wg.Done()
	file := FileInfo{
		Path: path,
		Size: fileInfo.Size(),
	}
	filesChan <- file
}
