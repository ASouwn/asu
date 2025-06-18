package router

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/evanw/esbuild/pkg/api"
)

func RouterInit(dir string) {
	routers, layoutDP := dirSeek(dir)

	for path, layouts := range layoutDP {
		log.Println("Layout path:", path)
		log.Println("Layouts found:", layouts)
	}

	if len(routers) == 0 {
		panic("No router found")
	}

	// create router for each directory that contains a page.tsx file
	for _, router := range routers {
		pageDir := strings.Split(router, "/")
		routePath := "/"
		// because the router is a directory, we need to remove the last part of the path
		// just like the path ./src/app/ the path is src/app, the router is "/"
		// so we need to remove the first two parts of the path
		if len(pageDir) > 2 {
			routePath += strings.Join(pageDir[2:], "/")
		}
		log.Println("Route path:", routePath)
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
