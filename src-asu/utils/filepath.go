package utils

import (
	"log"
	"os"
	"path/filepath"
)

func BfsWalkDir(root string, walk func(path string, d os.DirEntry, err error) error) error {
	queue := []string{root}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		entrys, err := os.ReadDir(path)

		if err != nil {
			log.Println("Error reading [directory]:", path, err)
		}
		for _, d := range entrys {
			fullPath := filepath.Join(path, d.Name())
			if err := walk(fullPath, d, nil); err != nil {
				return err
			}
			// if directory, add to queue
			if d.IsDir() {
				queue = append(queue, fullPath)
			}
		}
	}
	return nil
}
