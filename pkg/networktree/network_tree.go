package networktree

import (
	"bytes"
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

// Find finds the given node whose network has the given CIDR
func (t *Tree) Find(net *net.IPNet) *Tree {
	if t.Value.IP.Equal(net.IP) && bytes.Equal(t.Value.Mask, net.Mask) {
		return t
	}

	if t.Left != nil {
		if val := t.Left.Find(net); val != nil {
			return val
		}
	}
	if t.Right != nil {
		if val := t.Right.Find(net); val != nil {
			return val
		}
	}
	return nil
}

// MarkUsed marks a tree as being used
func (t *Tree) MarkUsed() {
	// fmt.Printf("Marking %s as used\n", t.Value.String())
	t.InUse = true
	if t.Left != nil {
		t.Left.MarkUsed()
	}
	if t.Right != nil {
		t.Right.MarkUsed()
	}
}

// UnusedRanges returns a slice of all the unused networks that you can assign
// grouped into their largest network.
func (t *Tree) UnusedRanges() []*net.IPNet {
	ranges := make([]*net.IPNet, 0)
	if t.areAllChildrenUnused() {
		return append(ranges, t.Value)
	}

	if t.Left != nil {
		ranges = append(ranges, t.Left.UnusedRanges()...)
	}

	if t.Right != nil {
		ranges = append(ranges, t.Right.UnusedRanges()...)
	}

	return ranges
}

func (t *Tree) areAllChildrenUnused() bool {
	if t.Left == nil && t.Right == nil {
		return !t.InUse
	}
	return (t.Left.areAllChildrenUnused() && t.Right.areAllChildrenUnused())
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
