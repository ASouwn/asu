package utils

import (
	"io/fs"
	"os"
	"path/filepath"
)

// Rewritten version of filepath.WalkDir
// All paths will be converted to forward slash (/) format
func DFSWalkDir(root string, fn fs.WalkDirFunc) error {
	rootpath := filepath.ToSlash(filepath.Clean(root))
	info, err := os.Lstat(rootpath)
	if err != nil {
		err = fn(rootpath, nil, err)
	} else {
		err = dfsWalkDir(rootpath, fs.FileInfoToDirEntry(info), fn)
	}
	if err == filepath.SkipDir || err == filepath.SkipAll {
		return nil
	}
	return err
}

func FilepathJoin(elem ...string) string {
	path := filepath.Join(elem...)
	return filepath.ToSlash(path) // Convert to forward slashes
}
