package router

import (
	"log"
	"net/http"
	"strings"
)

// RouterInit
//
// Parameters:
//   - routers: the dirs which have file page.tsx
//   - render: the way how to render the page.tsx to html
//
// Notes:
//   - render should be choose between the func file src-asu/render/render.go declared
func RouterInit(routers []string, rootpathLen int, render func(string) string) {
	// lazy init
	rendedPage := make(map[string][]byte)
	for _, router := range routers {
		rendedPage[router] = []byte(render(router))
	}
	// create router for each directory that contains a page.tsx file
	for _, router := range routers {
		pageDir := strings.Split(router, "/")
		routePath := "/"
		// because the router is a directory, we need to remove the last part of the path
		// just like the path ./src/app/ the path is src/app, the router is "/"
		// so we need to remove the first two parts of the path
		if len(pageDir) > rootpathLen {
			routePath += strings.Join(pageDir[rootpathLen:], "/")
		}
		log.Println("Route path:", routePath)
		http.HandleFunc(routePath, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(rendedPage[router]))
		})
	}
}
