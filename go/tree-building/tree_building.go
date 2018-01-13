package tree

import (
	"errors"
)

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {

	n := len(records)
	if n == 0 {
		return nil, nil
	}

	tree := make([]Node, len(records))
	sorted, err := SortRecords(records)

	if err != nil {
		return nil, err
	}

	for _, r := range sorted[1:] {
		tree[r.ID].ID = r.ID
		tree[r.Parent].Children = append(tree[r.Parent].Children, &tree[r.ID])
	}

	return &tree[0], nil

}

func SortRecords(records []Record) ([]Record, error) {
	n := len(records)
	sorted := make([]Record, n)

	var err error
	for _, r := range records {
		err = errChk(r, n)
		if err != nil {
			return nil, err
		}

		sorted[r.ID] = r
	}

	return sorted, nil
}

func errChk(r Record, n int) error {

	if r.ID == 0 && r.Parent == 0 {
		return nil
	} else if r.Parent >= r.ID {
		return errors.New("parent ID >= child ID")
	} else if r.ID >= n {
		return errors.New("non-continuous")
	} else {
		return nil
	}

}
