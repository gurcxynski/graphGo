package graph

type Graph struct {
	vertices []vertex
	free_id  int
}

func (src *Graph) Add(value int) int {
	new := vertex{src.free_id, value, []edge{}}
	src.vertices = append(src.vertices, new)
	src.free_id += 1
	return src.free_id - 1
}

func (src *Graph) Remove(id int) bool {
	id = src.find(id)
	if id == -1 {
		return false
	}
	for _, vertex := range src.vertices {
		for _, edge := range vertex.edges {
			if edge.target == id {
				vertex.unlink(id)
			}
			src.vertices[vertex.id] = vertex
		}
	}
	src.vertices[id] = src.vertices[len(src.vertices)-1]
	src.vertices = src.vertices[:len(src.vertices)-1]
	return true
}

func (src *Graph) Link(a, b, value int) bool {
	a = src.find(a)
	b = src.find(b)
	if a == -1 || b == -1 {
		return false
	}
	src.vertices[a].link(b, value)
	src.vertices[b].link(a, value)
	return true
}

func (src *Graph) Unlink(a, b int) bool {
	a1 := src.find(a)
	b1 := src.find(b)
	if a1 == -1 || b1 == -1 {
		return false
	}
	src.vertices[a1].unlink(b)
	src.vertices[b1].unlink(a)
	return true
}

func (src *Graph) find(id int) int {
	for i, v := range src.vertices {
		if v.id == id {
			return i
		}
	}
	return -1
}
