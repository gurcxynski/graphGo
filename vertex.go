package graph

type vertex struct {
	id    int
	value int
	edges []edge
}

func (src *vertex) link(target, value int) edge {
	new := edge{target, value}
	src.edges = append(src.edges, new)
	return new
}

func (src *vertex) unlink(id int) bool {
	id = src.find(id)
	src.edges[id] = src.edges[len(src.edges)-1]
	src.edges = src.edges[:len(src.edges)-1]
	return true
}

func (src *vertex) find(id int) int {
	for i, v := range src.edges {
		if v.target == id {
			return i
		}
	}
	return -1
}
