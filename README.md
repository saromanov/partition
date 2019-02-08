# partition
[![GoDoc](https://godoc.org/github.com/saromanov/partition?status.png)](https://godoc.org/github.com/saromanov/partition)
[![Go Report Card](https://goreportcard.com/badge/github.com/saromanov/partition)](https://goreportcard.com/report/github.com/saromanov/partition)
[![Build Status](https://travis-ci.org/saromanov/partition.svg?branch=master)](https://travis-ci.org/saromanov/partition)
[![Coverage Status](https://coveralls.io/repos/github/saromanov/partition/badge.svg?branch=master)](https://coveralls.io/github/saromanov/partition?branch=master)


Partition of collection on chunks

### Example
```go
package main

import (
	"fmt"

	"github.com/saromanov/partition"
)

func main() {
	value := []int{5, 2, 8, 9, 5, 3, 2, 0, 0, 2, 5, 5}
	for idx := range gopart.Partition(len(value), 2) {
		fmt.Println(value[idx.Start:idx.End])
	}
}
```

```
[5 2]
[8 9]
[5 3]
[2 0]
[0 2]
[5 5]
```
