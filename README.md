# Go int set microbenchmark

Should we implement a set of uint as sorted []uint with binary search or map[uint]struct{} ?

```sh
go test -run x -bench BenchmarkUidSets -benchtime 100000000x
```

That test run on AMD Ryzen 7 9700X

| set size | array | map |
| --- | ---- | ---- |
| 10 | 3.952 | 5.911 |
| 100 | 5.843 | 7.104 |
| 1000 | 7.899 | 6.008 |
| 10000 | 10.35 | 6.746 |
| 100000 | 13.15 | 9.547 |

`slices.BinarySearch` is faster below 1000 elements in the set. Somewhere between 100 and 1000 elements in the set, lookup into `map[uint]struct{}` becomes faster.
