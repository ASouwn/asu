package srcasu

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/ASouwn/asu/src-asu/render"
	"github.com/ASouwn/asu/src-asu/router"
	"github.com/ASouwn/asu/src-asu/utils"
)

// ASUStart start the server of asu.
//
// asu can package the page.tsx to html by ssr, and the rout is the dir where the page.tsx is
//
// Params:
//   - dir: the frontend files root dir, also the root rout start at.
//     like input ./src/app, the dir ./src/app/main with page.tsx will be found on rout /main
//   - port: the port the frontend server start at.
func ASUStart(dir, port string) {
	rootdir := filepath.ToSlash(filepath.Clean(dir))
	// Scan the frontend directory and return directories containing page.tsx and layout.tsx
	routers, _ := utils.DirSeek(rootdir)
	if len(routers) == 0 {
		panic("No router found")
	}
	// Render and route directories containing page.tsx
	router.RouterInit(routers, len(strings.Split(rootdir, "/")), render.SSRRender)

	fmt.Printf("Starting server on port http://localhost:%s ...\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
