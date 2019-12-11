package ordered_map

type OrderedMap struct {
	m map[interface{}]*orderedMapNode
}

type orderedMapNode struct {
	prev, next, hnext *orderedMapNode
}
