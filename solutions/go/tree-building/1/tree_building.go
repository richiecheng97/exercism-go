package tree

import "sort"

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	var tree *Node

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	for _, v := range records {
		node := &Node{ID: v.ID}
		if v.ID == 0 {
			tree = node
			continue
		}
		parent := findNode(tree, v.Parent)
		if parent != nil {
			parent.Children = append(parent.Children, node)
		}
	}

	return tree, nil
}

func findNode(node *Node, id int) *Node {
	if node.ID == id {
		return node
	}

	for _, child := range node.Children {
		if found := findNode(child, id); found != nil {
			return found
		}
	}
	return nil
}
