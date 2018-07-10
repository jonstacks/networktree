package networktree

import (
	"net"
)

// Tree is a network tree
type Tree struct {
	Left  *Tree
	Right *Tree
	Value *net.IPNet
	InUse bool
}

// New creates a new network tree
func New(cidr string) (*Tree, error) {
	_, n, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}
	if networkSize(n) <= 1 {
		return &Tree{nil, nil, n, false}, nil
	}

	s1, s2 := SplitNetwork(n)
	var left, right *Tree

	if s1 != nil {
		left, err = New(s1.String())
		if err != nil {
			return nil, err
		}
	}

	if s2 != nil {
		right, err = New(s2.String())
		if err != nil {
			return nil, err
		}
	}

	return &Tree{
		left,
		right,
		n,
		false,
	}, nil
}

// Height returns the height of the networktree
func (t *Tree) Height() int {
	if t.Left == nil && t.Right == nil {
		return 1
	}
	return 1 + max(t.Left.Height(), t.Right.Height())
}

func networkSize(n *net.IPNet) int {
	return 1 << unmaskSize(n.Mask)
}

// returns the number of ones in the mask
func unmaskSize(m net.IPMask) uint {
	ones, bits := m.Size()
	return uint(bits - ones)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
