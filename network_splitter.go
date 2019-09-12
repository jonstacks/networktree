package networktree

import (
	"math/big"
	"net"
)

// SplitNetwork splits a network into 2 subnets.
func SplitNetwork(n *net.IPNet) (*net.IPNet, *net.IPNet) {
	if networkSize(n) <= 1 {
		return nil, nil
	}

	newMask := incMask(n.Mask)
	s1 := &net.IPNet{
		IP:   n.IP,
		Mask: newMask,
	}
	xod := 1 << unmaskSize(newMask)
	ip2 := new(big.Int).SetBytes(n.IP)
	ip2.Xor(ip2, big.NewInt(int64(xod)))

	s2 := &net.IPNet{
		IP:   ip2.Bytes(),
		Mask: newMask,
	}
	return s1, s2
}

// incMask increases the size of the mask by 1 unless it is already a broadcast
// mask
func incMask(m net.IPMask) net.IPMask {
	ones, bits := m.Size()
	if ones >= bits {
		return m
	}
	return net.CIDRMask(ones+1, bits)
}
