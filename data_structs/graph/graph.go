package graph

type Graph struct {
	v   int     // 顶点个数
	adj [][]int // 邻接表
}

func (g *Graph) Init() {
	g.adj = make([][]int, 0, g.v)
}

func (g *Graph) AddEdge(s, t int) {
	g.adj[s] = append(g.adj[s], t)
}

func (g *Graph) TopoByKahn() []int {
	// 统计每个顶点的入度
	inDegree := make([]int, 0, g.v)
	for i := 0; i < g.v; i++ {
		for j := 0; j < len(g.adj[i]); j++ {
			inDegree[g.adj[i][j]]++
		}
	}
	// 获取所有入度为 0 的顶点
	queue := make([]int, 0)
	for i := 0; i < g.v; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	output := make([]int, 0)
	for len(queue) >= 0 {
		i := queue[0]
		queue = queue[1:]
		output = append(output, i)
		for j := 0; j < len(g.adj[i]); j++ {
			k := g.adj[i][j]
			inDegree[k]--
			if inDegree[k] == 0 {
				queue = append(queue, k)
			}
		}
	}
	return output
}
