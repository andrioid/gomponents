package gomponents

import (
	"fmt"
)

func concatAttrs(nodes []Node) []Node {
	byName := make(map[string]int)
	result := nodes[:0]
	for _, n := range nodes {
		a, ok := n.(*attr)
		if !ok {
			result = append(result, n)
			continue
		}

		if j, ok := byName[a.name]; ok {
			c := fmt.Sprintf("%s %s", *nodes[j].(*attr).value, *a.value)
			result[j] = &attr{
				name:  a.name,
				value: &c,
			}
			continue

		}

		byName[a.name] = len(result)
		result = append(result, n)
	}
	return result
}

// flattenGroups recursively flattens groups in children.
func flattenGroups(children []Node) []Node {
	var flat []Node
	for _, c := range children {
		if g, ok := c.(Group); ok {
			flat = append(flat, flattenGroups(g)...)
		} else {
			flat = append(flat, c)
		}
	}
	return flat
}
