package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func (n *Node) AddRecord(r Record, processed map[int]*Node) {
	newNode := &Node{ID: r.ID}
	n.Children = append(n.Children, newNode)
	processed[r.ID] = newNode
}

func SortRecords(records []Record) {

	sort.Slice(records[:], func(i, j int) bool {
		if records[i].Parent == records[j].Parent {
			return records[i].ID < records[j].ID
		}
		return records[i].Parent < records[j].Parent
	})
}

func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}

	SortRecords(records)

	if records[0].ID != 0 {
		return nil, errors.New("no root")
	} else if records[0].Parent != 0 {
		return nil, errors.New("root has parent")
	}

	rootID := records[0].ID
	root := &Node{ID: rootID}
	processed := make(map[int]*Node)
	processed[rootID] = root
	current := processed[rootID]

	for _, r := range records[1:] {

		err := errChk(records, processed, r)
		if err != nil {
			return nil, err
		}

		if current.ID != r.Parent {
			current = processed[r.Parent]
		}

		current.AddRecord(r, processed)
	}

	return root, nil
}

func errChk(records []Record, processed map[int]*Node, r Record) error {

	if processed[r.ID] != nil {
		return errors.New("duplicate node")
	} else if r.ID >= len(records) {
		return errors.New("non-continuous")
	} else if r.Parent > r.ID {
		return errors.New("parent ID > child ID")
	} else if processed[r.Parent] == nil {
		return errors.New("detached tree")
	} else {
		return nil
	}

}
