package main

type node [10]bool

type edge struct {
	from   node
	button []uint
	to     node
}

type graph struct {
	nodes map[node][]edge
}

func newGraph() *graph {
	g := &graph{
		nodes: make(map[node][]edge),
	}

	return g
}

func (g *graph) addEdge(from node, button []uint, to node) {
	e := edge{
		from:   from,
		button: button,
		to:     to,
	}

	g.nodes[from] = append(g.nodes[from], e)
}
