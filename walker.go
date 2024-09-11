package main

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"
)

func walkDir(dir string, filesChan chan FileInfo, wg *sync.WaitGroup) {
	defer wg.Done()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("Error reading directory %s: %v", dir, err)
		return
	}

	for _, entry := range entries {
		fullPath := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			wg.Add(1)
			go walkDir(fullPath, filesChan, wg)
		} else {
			wg.Add(1)
			go processFile(fullPath, entry, filesChan, wg)
		}
	}
}
