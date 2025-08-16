package srcasu

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ASouwn/asu/src-asu/types"
	"github.com/ASouwn/asu/src-asu/utils"
)

func TestDirseek(t *testing.T) {
	err := os.Chdir("../")
	if err != nil {
		log.Fatal(err)
	}
	routers, dpmap := utils.DirSeek("src/app")
	fmt.Println("Routers:", routers)
	fmt.Println("Dynamic Params Map:")
	for str := range dpmap {
		fmt.Printf("%s : %s\n", str, dpmap[str])
	}
}

func TestInitTree(t *testing.T) {
	err := os.Chdir("../")
	if err != nil {
		log.Fatal(err)
	}
	routeTree, pagemap, err := types.NewRouteNode("src/app")
	if err != nil {
		t.Fatalf("Failed to initialize route tree: %v", err)
	}
	if routeTree == nil {
		t.Fatal("Route tree is nil")
	}
	fmt.Println("Route Tree Initialized Successfully")
	for path, node := range pagemap {
		fmt.Printf("Path: %s, Layout: %t, Page: %t\n", path, node.Layout, node.Page)
	}
}
