# networktree

[![codecov](https://codecov.io/gh/jonstacks/networktree/branch/master/graph/badge.svg)](https://codecov.io/gh/jonstacks/networktree)
[![CI](https://github.com/jonstacks/networktree/actions/workflows/ci.yml/badge.svg)](https://github.com/jonstacks/networktree/actions/workflows/ci.yml)
[![golangci](https://github.com/jonstacks/networktree/actions/workflows/golangci.yml/badge.svg)](https://github.com/jonstacks/networktree/actions/workflows/golangci.yml)
[![GoDoc](https://godoc.org/github.com/jonstacks/networktree?status.png)](https://godoc.org/github.com/jonstacks/networktree)

A simple library for calculating the biggest unused subnets in a network.

[Example](https://github.com/jonstacks/aws/blob/42e4c0c207572badc28b974f27c6fca233ac66d1/pkg/views/ec2.go#L277-L299) usage with AWS SDK to find unused ranges in a VPC:

```go

		cidr := aws.StringValue(vpc.CidrBlock)
		tree, err := networktree.New(cidr)
		if err != nil {
			fmt.Fprintf(w, "%s", err)
			continue
		}
		for _, subnet := range subnets {
			subCidr := aws.StringValue(subnet.CidrBlock)
			_, sn, err := net.ParseCIDR(subCidr)
			if err != nil {
				continue
			}
			subtree := tree.Find(sn)
			if subtree != nil {
				subtree.MarkUsed()
			}
		}

		unusedRanges := tree.UnusedRanges()
		unusedCIDRs := make([]string, len(unusedRanges))
		for i, n := range unusedRanges {
			unusedCIDRs[i] = n.String()
		}
```

*Note:* I have not done thourough performance testing on this yet.
