package gopart

// Range defines value for split collection
type Range struct {
	Start, End int
}

// Partition returns index after partition
func Partition(size, part int) chan Range {
	c := make(chan Range)
	if part <= 0 {
		close(c)
		return c
	}

	go func() {
		numFullPartitions := size / part
		var i int
		for ; i < numFullPartitions; i++ {
			high := (i + 1) * part
			c <- Range{Start: i * part, End: high}
		}

		if size%part != 0 {
			c <- Range{Start: i * part, End: size}
		}

		close(c)
	}()
	return c
}
