package tree

import (
	"errors"
	"sort"
)

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
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	// Basic root validation
	if records[0].ID != 0 || records[0].Parent != 0 {
		return nil, errors.New("invalid root")
	}

	// IDs must be continuous from 0..n-1 and parents must be valid (< id)
	for i, r := range records {
		if r.ID != i {
			return nil, errors.New("non-continuous or duplicate id")
		}
		if r.ID == 0 && r.Parent != 0 {
			return nil, errors.New("root has parent")
		}
		if r.ID != 0 {
			// parent must be less than the node id
			if r.Parent < 0 || r.Parent >= r.ID {
				return nil, errors.New("invalid parent")
			}
		}
	}

	// Build nodes indexed by ID for O(1) parent lookup
	nodes := make([]*Node, len(records))
	var root *Node
	for _, r := range records {
		node := &Node{ID: r.ID}
		nodes[r.ID] = node
		if r.ID == 0 {
			root = node
			continue
		}
		parent := nodes[r.Parent]
		if parent == nil {
			// Shouldn't happen because of previous checks, but guard anyway
			return nil, errors.New("parent not found")
		}
		parent.Children = append(parent.Children, node)
	}

	return root, nil
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
