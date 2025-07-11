package srcasu

import (
	"fmt"
	"log"
	"os"
	"testing"

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
