package intsets

import (
	"fmt"
	"slices"
)

type UintSet interface {
	Set(uid uint)
	Contains(uid uint) bool
	Clear()
}

type ArrayUintSet struct {
	they []uint
}

func (a *ArrayUintSet) Clear() {
	if a.they != nil {
		a.they = a.they[:0]
	}
}

func copyUintSliceWithInsert(uids []uint, newUid uint) (newUids []uint, alreadyThere bool) {
	insertPos, found := slices.BinarySearch(uids, newUid)
	if found {
		return uids, true
	}
	out := make([]uint, len(uids)+1)
	if insertPos > 0 {
		copy(out[:insertPos], uids[:insertPos])
	}
	out[insertPos] = newUid
	if insertPos < len(uids) {
		copy(out[insertPos+1:], uids[insertPos:])
	}
	return out, false
}

func (a *ArrayUintSet) Set(uid uint) {
	a.they, _ = copyUintSliceWithInsert(a.they, uid)
}

func (a *ArrayUintSet) Contains(uid uint) bool {
	_, found := slices.BinarySearch(a.they, uid)
	return found
}

func (a *ArrayUintSet) String() string {
	return fmt.Sprintf("%v", a.they)
}

type MapUintSet struct {
	they map[uint]struct{}
}

func (m *MapUintSet) Clear() {
	m.they = make(map[uint]struct{})
}

func (m *MapUintSet) Set(uid uint) {
	m.they[uid] = struct{}{}
}
func (m *MapUintSet) Contains(uid uint) bool {
	_, found := m.they[uid]
	return found
}

func (m *MapUintSet) String() string {
	return fmt.Sprintf("%v", m.they)
}
