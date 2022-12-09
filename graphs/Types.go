package graphs

type Edge struct {
	Weight int
	Start  Node
	End    Node
}

type Node struct {
	Name string
}

type WeightedNode struct {
	Parent   *WeightedNode
	Children []*WeightedNode
	Weight   int
}
