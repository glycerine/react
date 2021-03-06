// Code generated by immutableGen. DO NOT EDIT.

// Copyright (c) 2016 Paul Jolly <paul@myitcv.org.uk>, all rights reserved.
// Use of this document is governed by a license found in the LICENSE document.

package example

//go:generate echo "hello world"
//immutableVet:skipFile

import (
	"myitcv.io/immutable"
)

//
// myTestMap is an immutable type and has the following template:
//
// 	map[string]int
//
type myTestMap struct {
	theMap  map[string]int
	mutable bool
	__tmpl  _Imm_myTestMap
}

var _ immutable.Immutable = new(myTestMap)
var _ = new(myTestMap).__tmpl

func newMyTestMap(inits ...func(m *myTestMap)) *myTestMap {
	res := newMyTestMapCap(0)
	if len(inits) == 0 {
		return res
	}

	return res.WithMutable(func(m *myTestMap) {
		for _, i := range inits {
			i(m)
		}
	})
}

func newMyTestMapCap(l int) *myTestMap {
	return &myTestMap{
		theMap: make(map[string]int, l),
	}
}

func (m *myTestMap) Mutable() bool {
	return m.mutable
}

func (m *myTestMap) Len() int {
	if m == nil {
		return 0
	}

	return len(m.theMap)
}

func (m *myTestMap) Get(k string) (int, bool) {
	v, ok := m.theMap[k]
	return v, ok
}

func (m *myTestMap) AsMutable() *myTestMap {
	if m == nil {
		return nil
	}

	if m.Mutable() {
		return m
	}

	res := m.dup()
	res.mutable = true

	return res
}

func (m *myTestMap) dup() *myTestMap {
	resMap := make(map[string]int, len(m.theMap))

	for k := range m.theMap {
		resMap[k] = m.theMap[k]
	}

	res := &myTestMap{
		theMap: resMap,
	}

	return res
}

func (m *myTestMap) AsImmutable(v *myTestMap) *myTestMap {
	if m == nil {
		return nil
	}

	if v == m {
		return m
	}

	m.mutable = false
	return m
}

func (m *myTestMap) Range() map[string]int {
	if m == nil {
		return nil
	}

	return m.theMap
}

func (m *myTestMap) WithMutable(f func(mi *myTestMap)) *myTestMap {
	res := m.AsMutable()
	f(res)
	res = res.AsImmutable(m)

	return res
}

func (m *myTestMap) WithImmutable(f func(mi *myTestMap)) *myTestMap {
	prev := m.mutable
	m.mutable = false
	f(m)
	m.mutable = prev

	return m
}

func (m *myTestMap) Set(k string, v int) *myTestMap {
	if m.mutable {
		m.theMap[k] = v
		return m
	}

	res := m.dup()
	res.theMap[k] = v

	return res
}

func (m *myTestMap) Del(k string) *myTestMap {
	if _, ok := m.theMap[k]; !ok {
		return m
	}

	if m.mutable {
		delete(m.theMap, k)
		return m
	}

	res := m.dup()
	delete(res.theMap, k)

	return res
}
func (s *myTestMap) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	if s == nil {
		return true
	}

	if s.Mutable() {
		return false
	}
	return true
}
