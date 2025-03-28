# Go int set microbenchmark

Should we implement a set of uint as sorted []uint with binary search or map[uint]struct{} ?

On my system, small sets are faster with slices.BinarySearch over a sorted []uint.
Between 100 and 1000 that flips and map is faster.

```sh
go test -run x -bench BenchmarkUidSets -benchtime 100000000x
```
