package intsets

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

// BenchmarkUintSets
// go test -v -bench . -benchtime 100000000x
func BenchmarkUintSets(b *testing.B) {
	setsizes := []int{10, 100, 1000, 10_000, 100_000}
	impls := []UintSet{&ArrayUintSet{}, &MapUintSet{they: make(map[uint]struct{})}}

	uidsByN := make(map[int][]uint)
	accessByN := make(map[int][]uint)

	setupN := func(bN, setsize int) (ouids, oaccess []uint) {
		ouids, ok := uidsByN[bN]
		oaccess, ok2 := accessByN[bN]
		if ok && ok2 {
			return
		}
		uids := make([]uint, setsize)
		for i := range uids {
			uids[i] = uint(i + 1)
		}
		access := make([]uint, bN)
		for i := range access {
			// 10% hit rate
			access[i] = uint(rand.UintN(uint(setsize * 10)))
		}
		uidsByN[bN] = uids
		accessByN[bN] = access
		return uids, access
	}

	for _, setsize := range setsizes {
		uidsByN = make(map[int][]uint)
		accessByN = make(map[int][]uint)

		for _, impl := range impls {
			implname := fmt.Sprintf("%T_%d", impl, setsize)
			b.Run(implname, func(b *testing.B) {
				impl.Clear()
				uids, access := setupN(b.N, setsize)
				//b.Logf("b.N %d, len(uids) %d, len(access) %d", b.N, len(uids), len(access))
				for _, xuid := range uids {
					impl.Set(xuid)
				}
				//b.Logf("impl %s", impl)
				b.ResetTimer()
				hitcount := uint64(0)
				const mult = 1
				for multi := 0; multi < mult; multi++ {
					for _, xuid := range access {
						has := impl.Contains(xuid)
						if has {
							hitcount++
						}
					}
				}
				if b.N >= 10_000 {
					expected := mult * 0.1 * float64(len(access))
					ratio := float64(hitcount) / expected
					if ratio < 0.90 {
						b.Errorf("hit rate low, %0.9f, hitcount %d, expected %f, impl %s", ratio, hitcount, expected, impl)
					}
					if ratio > 1.10 {
						b.Errorf("hit rate high, %0.9f", ratio)
					}
				}
			})
		}
	}
}
