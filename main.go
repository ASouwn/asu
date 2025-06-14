package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ASouwn/asu/src-asu/router"
)

func main() {
	port := "3000"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	/**
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result := api.Build(api.BuildOptions{
			EntryPoints: []string{"src/app/layout.tsx"},
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
	**/
	router.RouterInit("src/")
	fmt.Printf("Starting server on port http://localhost:%s ...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
