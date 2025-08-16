package types

import (
	"io/fs"
	"path/filepath"

	"github.com/ASouwn/asu/src-asu/utils"
)

func NewRouteNode(root string) (*RouteNode, map[string]*RouteNode, error) {
	rootpath, hashmap, pageMap := utils.FilepathFormat(root), make(map[string]*RouteNode), make(map[string]*RouteNode)
	err := utils.DFSWalkDir(rootpath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			nowpath := utils.FilepathFormat(path)
			node := &RouteNode{
				Path:     nowpath,
				Layout:   false,
				Page:     false,
				Parent:   nil,
				Children: []*RouteNode{},
			}
			hashmap[nowpath] = node
			if utils.FileExists(utils.FilepathJoin(nowpath, "layout.tsx")) {
				node.Layout = true
			}
			if utils.FileExists(utils.FilepathJoin(nowpath, "page.tsx")) {
				node.Page = true
				pageMap[nowpath] = node
			}
			if nowpath != rootpath {
				// Set parent node if not root
				if parentNode, exists := hashmap[utils.FilepathFormat(filepath.Dir(nowpath))]; exists {
					node.Parent = parentNode
					parentNode.Children = append(parentNode.Children, node)
				} else {
					panic("Parent node not found for " + nowpath)
				}
			}
		}
		return nil
	})
	if err != nil {
		panic("Error walking directory when init RouteTree: " + err.Error())
	}
	if _, exists := hashmap[rootpath]; !exists {
		panic("Root node not found in RouteTree")
	}
	return hashmap[rootpath], pageMap, nil
}

func NewRouteBuild(root string) (*RouteBuild, error) {
	tree, pagemap, _ := NewRouteNode(root)
	return &RouteBuild{
		RouteTree:   tree,
		PageNodeMap: pagemap,
	}, nil
}
