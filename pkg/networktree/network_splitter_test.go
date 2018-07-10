package networktree

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitNetwork(t *testing.T) {
	splitNets := map[string][2]string{
		"192.168.1.0/24":  [2]string{"192.168.1.0/25", "192.168.1.128/25"},
		"10.10.0.0/16":    [2]string{"10.10.0.0/17", "10.10.128.0/17"},
		"0.0.0.0/0":       [2]string{"0.0.0.0/1", "128.0.0.0/1"},
		"192.168.1.30/24": [2]string{"192.168.1.0/25", "192.168.1.128/25"},
	}

	for cidr, subnets := range splitNets {
		_, n, err := net.ParseCIDR(cidr)
		assert.Nil(t, err)
		s1, s2 := SplitNetwork(n)
		assert.Equal(t, subnets[0], s1.String())
		assert.Equal(t, subnets[1], s2.String())
	}
}
