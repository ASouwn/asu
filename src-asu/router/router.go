package router

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/evanw/esbuild/pkg/api"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}

func dirSeek(dir string) (routers []string) {
	// dfs for router files
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			pagePath := filepath.Join(path, "page.tsx")
			if fileExists(pagePath) {
				routers = append(routers, path)
			}
		}
		return nil
	})
	if err != nil {
		panic("Error walking directory: " + err.Error())
	}

	return routers
}

func RouterInit(dir string) {
	routers := dirSeek(dir)
	if len(routers) == 0 {
		panic("No router found")
	}

	for _, router := range routers {
		pageDir := strings.Split(router, "\\")
		log.Println("Router path:", router)
		// root router, its path is ../../src/app
		// so the length is 4
		routePath := "/"
		if len(pageDir) > 2 {
			routePath += strings.Join(pageDir[2:], "/")
		}
		http.HandleFunc(routePath, func(w http.ResponseWriter, r *http.Request) {
			result := api.Build(api.BuildOptions{
				EntryPoints: []string{filepath.Join(router, "layout.tsx")},
				Bundle:      true,
				Write:       false,
				Platform:    api.PlatformNode,
				Format:      api.FormatCommonJS,
				Loader: map[string]api.Loader{
					".tsx": api.LoaderTSX,
					".ts":  api.LoaderTS,
					".js":  api.LoaderJS,
				},
				External: []string{"react", "react-dom", "react-dom/server"},
			})
			cmd := exec.Command("node", "ssr.js")

			var stdout, stderr bytes.Buffer
			cmd.Stdin = bytes.NewReader(result.OutputFiles[0].Contents)
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			err := cmd.Run()
			if err != nil {
				http.Error(w, fmt.Sprintf("SSR error: %v\n%s", err, stderr.String()), 500)
				return
			}

			html := `<!DOCTYPE html><html><body>` + stdout.String() + `</body></html>`
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(html))
		})
	}
}
