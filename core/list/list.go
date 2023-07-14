package list

import "github.com/karlsing/data-structure/core/iterator"

type List[T any] interface {
	Length() int
	Get(index int) T
	Set(index int, elem T)
	Push(T)
	Concat(List[T]) List[T]
	Copy() List[T]
	Sub(start, length, step int) List[T]
	ToSlice() []T
	Iter() iterator.Iterable[T]
}

// Find returns the index of param 'target' or -1 if not found in list
func Find[T comparable](list List[T], target T) int {
	index := -1
	iterator.ForeachB(list.Iter(), func(i int, elem T) bool {
		if elem == target {
			index = i
			return true
		}
		return false
	})
	return index
}

// Contains return if 'target' was contained by 'list'
func Contains[T comparable](list List[T], target T) bool {
	return Find(list, target) >= 0
}
