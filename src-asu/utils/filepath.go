package utils

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
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

// dirSeek traverses the directory structure starting from the given directory
// and returns a slice of strings containing the paths of directories that
// contain a "page.tsx" file.
//
// Parameters:
//   - dir: the root directory to start searching from
//
// Returns:
//   - routers: the paths of directories containing "page.tsx"
//   - layoutDP: the sorted paths of directories containing "layout.tsx" by the directory depth
func DirSeek(dir string) (routers []string, layoutDP map[string][]string) {
	// Initialize the routers slice and layoutDP map
	layoutDP = make(map[string][]string)
	// dfs for router files
	err := DFSWalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			// log.Println("try to build router and layout tree for [path]-", path)

			pagePath := filepath.Join(path, "page.tsx")
			// log.Println("pagePath:", pagePath)
			if FileExists(pagePath) {
				routers = append(routers, path)
			}

			layoutPath := filepath.Join(path, "layout.tsx")
			dir := strings.Split(path, "/")
			// log.Println("parent dir:", strings.Join(dir[:len(dir)-1], "/"))
			if parentSlice, ok := layoutDP[strings.Join(dir[:len(dir)-1], "/")]; ok {
				// 创建一个新的切片拷贝
				layoutDP[path] = append([]string{}, parentSlice...)
			} else {
				layoutDP[path] = []string{}
			}
			if FileExists(layoutPath) {
				layoutDP[path] = append(layoutDP[path], path)
			}
		}
		return nil
	})
	if err != nil {
		panic("Error walking directory: " + err.Error())
	}

	return routers, layoutDP
}

func FilepathFormat(path string) string {
	return filepath.ToSlash(filepath.Clean(path))
}

// FileExists checks if a file exists at the given path
//
// Returns:
//   - true if the file exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}
