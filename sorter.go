package main

import (
	"sort"
)

func sortFilesBySize(files []FileInfo) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Size > files[j].Size
	})
}
