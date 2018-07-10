package networktree

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNetworkSize(t *testing.T) {
	netSizes := map[string]int{
		"192.168.0.0/24": 256,
		"0.0.0.0/0":      4294967296,
		"10.10.10.10/32": 1,
		"192.168.0.0/18": 16384,
		"10.10.10.10/16": 65536,
	}

	for cidr, size := range netSizes {
		_, n, err := net.ParseCIDR(cidr)
		assert.Nil(t, err)
		assert.Equal(t, size, networkSize(n))
	}
}

func TestTreeHeight(t *testing.T) {
	trees := map[string]int{
		"192.168.1.0/24":   9,
		"10.10.0.0/16":     17,
		"192.168.1.196/32": 1,
		"192.169.1.0/8":    25,
	}

	for cidr, height := range trees {
		tree, err := New(cidr)
		assert.Nil(t, err)
		assert.Equal(t, height, tree.Height())
	}
}

func TestTreeFind(t *testing.T) {
	results := map[string]bool{
		"192.168.0.0/24":   true,
		"192.168.0.128/25": true,
		"192.128.0.0/24":   false,
	}
	tree, err := New("192.168.0.0/16")
	assert.Nil(t, err)

	for cidr, shouldFind := range results {
		_, n, err := net.ParseCIDR(cidr)
		assert.Nil(t, err)

		if shouldFind {
			assert.NotNil(t, tree.Find(n))
		} else {
			assert.Nil(t, tree.Find(n))
		}
	}

}
