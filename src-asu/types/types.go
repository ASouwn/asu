package types

type RouteNode struct {
	Path     string
	Layout   bool
	Page     bool
	Parent   *RouteNode
	Children []*RouteNode
}

type RouteBuild struct {
	RouteTree   *RouteNode
	PageNodeMap map[string]*RouteNode
}
