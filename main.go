package main

import (
	"os"

	srcasu "github.com/ASouwn/asu/src-asu"
)

func main() {
	port := "3000"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	// the input dir will decide the root rout "/"
	// its ok to use ./path or /path, even ./path/../path2 to replace /path2
	// also "/" and "\\" will be fine
	srcasu.ASUStart("./src/app", port)
}
