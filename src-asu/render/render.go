package render

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/ASouwn/asu/src-asu/utils"
	"github.com/evanw/esbuild/pkg/api"
)

// Render a directory containing a page.tsx file
//
// Parameters:
//   - router: the dir has page.tsx
//
// Returns:
//   - rende the page.tsx to html string
//
// Note:
//   - Currently, only layout.tsx under the router directory can be rendered directly
func SSRRender(router string) string {
	result := api.Build(api.BuildOptions{
		EntryPoints: []string{utils.FilepathJoin(router, "layout.tsx")},
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
		return fmt.Sprintf("SSR error: %v\n%s", err, stderr.String())
	}
	return `<!DOCTYPE html><html><body>` + stdout.String() + `</body></html>`
}
