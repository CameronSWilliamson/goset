package goset

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type OrderedComparable interface {
	comparable
	constraints.Ordered
}

type Set[T OrderedComparable] map[T]struct{}

func (s Set[T]) Add(item T) {
	s[item] = struct{}{}
}

func (s Set[T]) Remove(item T) {
	delete(s, item)
}

func FromList[T OrderedComparable](input []T) Set[T] {
	set := Set[T]{}
	for _, k := range input {
		set.Add(k)
	}
	return set
}

func (s Set[T]) ToList() []T {
	keys := make([]T, 0)
	for k := range s {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}
